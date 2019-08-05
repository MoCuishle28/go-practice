package scheduler

// 通过队列进行调度

import(

	"Go-practice/in-depth-study/reptile-project/engine"
)


// 要实现 engine 内定义的 Scheduler 接口
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request 	// chan 的类型是一个 engine.Request 的 chan
}


// 将 chan 传入
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {

}


func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}


func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}


func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ  []engine.Request
		var workerQ  []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select{
				case r := <-s.requestChan:
					requestQ = append(requestQ, r)
				case w := <-s.workerChan:
					workerQ = append(workerQ, w)
				case activeWorker <- activeRequest:
					// 出队
					workerQ = workerQ[1:]
					requestQ = requestQ[1:]
			}
		}
	} ()
}