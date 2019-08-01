package engine


import(
	"log"

	"Go-practice/in-depth-study/reptile-project/fetcher"
)


func Run(seeds ...Request) {
	var requests []Request 			// 维护一个队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 出队一个来处理
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error " + "fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult :=  r.ParserFunc(body)
		// 加三个点 意味着将所有内容取出加进 requests
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}