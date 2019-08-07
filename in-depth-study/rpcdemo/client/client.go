package main


import(
	"net"
	"net/rpc/jsonrpc"
	"fmt"

	"Go-practice/in-depth-study/rpcdemo"
)


func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var result float64
	client := jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result = 0
	err = client.Call("DemoService.Add", rpcdemo.Args{20, 8}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	var retStr string
	err = client.Call("DemoService.Change", rpcdemo.Args2{"Hello!"}, &retStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(retStr)
	}
}