package scheduler

// 简单的调度器


import(

	"Go-practice/in-depth-study/reptile-project/engine"
)


// 实现 engine 内定义的 Scheduler 接口
type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}


func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	// 不用实现
}


func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	// 不这样会导致死锁（循环等待）
    go func() {
    	s.workerChan <- r
    } ()
}