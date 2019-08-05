package parser

import (
	// "fmt"
	"regexp"

	"Go-practice/in-depth-study/reptile-project/engine"
)


const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


func ParseCityList(contents []byte) engine.ParseResult {
	// [^>]* 只要不是 > 都匹配 （^是除了的意思）
	// 用括号把想要的内容挖出来 这里只拿了 url、城市
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	limit := 5

	for _, m := range matches {	
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
		// fmt.Printf("city:%s, URL:%s\n", m[2], m[1])

		// 控制一下 只爬取5个城市
		limit--
		if limit == 0 {
			break
		}
	}
	// fmt.Printf("Matches found: %d\n", len(matches))
	return result
}