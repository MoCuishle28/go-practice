package engine


import(

)


// 并发版 engine

type ConcurrentEngine struct {
	Scheduler Scheduler 	// 调度器 	
	WorkerCount int 		// 多少个 worker
	ItemChan chan interface{} 	// 传输要存储 Item 的 chan
}


// 调度器接口
type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request 		// 由 Scheduler 要 chan （不用自己判断是每人一个chan还是全部共用一个chan）
	Run()
}


type ReadNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			// 得到 Items 后尽快送出去
			go func() {
				e.ItemChan <- item
			} ()
		}

		// 把 item 的 Requests 送给调度器
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}


// 抽象出 worker ：包括 Parser 和 Fetcher (输入 Request 处理后输出 Requests、Items)
func createWorker(in chan Request, out chan ParseResult, ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	} ()
}