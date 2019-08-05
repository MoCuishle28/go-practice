package parser

import(
	"regexp"
	// "fmt"

	"Go-practice/in-depth-study/reptile-project/engine"
)


const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`


func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])	// 要先将 m[2] 拷贝出来，否则下面的函数只会等待调用时再使用引用的最后一个 m[2]
		// result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(m[1]),
			// 用函数式编程封装 不要修改 types.go 中解析函数的定义
			ParserFunc:	func(c []byte) engine.ParseResult{
				return ParserProfile(c, name)
			},
		})
		// fmt.Printf("Got item User:%s URL:%s\n", m[2], m[1])
	}
	return result
}