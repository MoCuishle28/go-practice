package main

import(
	"Go-practice/in-depth-study/reptile-project/engine"
	"Go-practice/in-depth-study/reptile-project/zhenai/parser"
	// "golang.org/x/text/encoding/simplifiedchinese"
)


func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}