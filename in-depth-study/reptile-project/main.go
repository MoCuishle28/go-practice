package main

import(
	"Go-practice/in-depth-study/reptile-project/engine"
	"Go-practice/in-depth-study/reptile-project/zhenai/parser"

	"Go-practice/in-depth-study/reptile-project/scheduler"
	"Go-practice/in-depth-study/reptile-project/persist"
	// "golang.org/x/text/encoding/simplifiedchinese"
)


func main() {
	// 调用单任务版 engine
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url: "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// 并发版
	e := engine.ConcurrentEngine{
		// 使用简单调度器
		// Scheduler: &scheduler.SimpleScheduler{},

		// 使用队列调度器
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}