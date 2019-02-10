package main
import (
	"net"
	"os"
	"fmt"
	"reflect"
	// "io/ioutil"
)

func main() {
	// 这样运行，183.232.231.172 就是[1]   go run socket-test.go 183.232.231.172
	fmt.Println(os.Args, reflect.TypeOf(os.Args), len(os.Args))	// os.Arg 似乎是cmd输入的...

	if len(os.Args) < 2 {
		// [0] 是文件路径 	[1] 是用户输入参数
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)	// 把一个IPv4或者IPv6的地址转化成IP类型
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	fmt.Println("-----conn 127.0.0.1 service------")
	service := os.Args[1]
	msg := os.Args[2]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)		// 转换为 TCPAddr
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)			// 建立TCP连接,获得 TCPConn
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte(msg))  // 发送 http 请求
	checkError(err)

	fmt.Println("start ReadAll")

	result := make([]byte, 256)
	// result, err := ioutil.ReadAll(conn)		// 读取返回内容
	result_len, err := conn.Read(result)
	checkError(err)

	fmt.Println(string(result))
	fmt.Println("len:", result_len)
	
	fmt.Println("end ReadAll")

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}