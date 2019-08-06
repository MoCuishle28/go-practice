package main


// RPC 服务器

import(
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"Go-practice/in-depth-study/rpcdemo"
)


func main() {
	// 注册提供 RPC 服务的结构体
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}