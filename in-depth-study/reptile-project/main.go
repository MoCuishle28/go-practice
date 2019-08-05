package main

import(
	"Go-practice/in-depth-study/reptile-project/engine"
	"Go-practice/in-depth-study/reptile-project/zhenai/parser"

	"Go-practice/in-depth-study/reptile-project/scheduler"
	
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
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}