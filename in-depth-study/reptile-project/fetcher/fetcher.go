package fetcher

import(
	"bufio"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)


var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	// 100毫秒收一个 防止请求过快
	<-rateLimiter

	// 设置请求头应对403错误
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("err:%v\n", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// 不成功请求的时候
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	// 根据网页编码转为utf8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}


// 获取网页编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	// 封装一下为了其他地方可以重复读 读取前面1024个字节
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v\n", err)
		return unicode.UTF8
	}

	// 猜 html 的 encoding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}