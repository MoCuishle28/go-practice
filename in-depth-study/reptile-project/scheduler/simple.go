package scheduler

// 简单的调度器


import(

	"Go-practice/in-depth-study/reptile-project/engine"
)


// 实现 engine 内定义的 Scheduler 接口
type SimpleScheduler struct {
	workerChan chan engine.Request
}


// 将 chan 传入
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	// 这样会导致死锁（循环等待）
    go func() {
    	s.workerChan <- r
    } ()
}