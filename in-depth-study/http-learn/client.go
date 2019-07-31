package main

import(
	"net/http"
	"net/http/httputil"
	"fmt"
)


func main() {
	// 设置请求头部 控制访问动作
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	// 设置请求头部 访问手机版的imooc
	request.Header.Add( "User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) "+
		"AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		// 查看是否重定向
		// 所有重定向路径放via 每次重定向目标放req
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
				fmt.Println("Redirect: ", req)
				fmt.Println("-----------------------")
				return nil
			},
	}

	// resp, err := http.DefaultClient.Do(request)			// 设置完头部要用这个发起请求, 也可以自己创建一个Client
	resp, err := client.Do(request) 					// 用自己创建的client发起请求
	// resp, err := http.Get("http://www.imooc.com")	// 不设置头部可以用这个发起请求
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)		// 返回一个 []byte (网页信息)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}