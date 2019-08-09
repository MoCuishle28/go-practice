package engine


import(
	"log"
)

// 单任务版 engine


type SimpleEngine struct{}


func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request 			// 维护一个队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 出队一个来处理
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// 加三个点 意味着将所有内容取出加进 requests
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}


