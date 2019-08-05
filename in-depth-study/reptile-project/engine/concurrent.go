package engine


import(
	"log"

	// "Go-practice/in-depth-study/reptile-project/fetcher"
)


// 并发版 engine

type ConcurrentEngine struct {
	Scheduler Scheduler 	// 调度器 	
	WorkerCount int 		// 多少个 worker
}


// 调度器接口
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}


func (e *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}

		// 把 item 的 Requests 送给调度器
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	} ()
}