package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	// 字段名对应JSON里面的KEY
	Servers []Server 	// 数组对应slice
}	

func main() {
	var s Serverslice
	str := `{
				"servers":
					[
						{
							"serverName":"Shanghai_VPN",
							"serverIP":"127.0.0.1"
						},
						{
							"serverName":"Beijing_VPN",
							"serverIP":"127.0.0.2"
						}
					]
			}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	fmt.Println("-------read json file--------")
	// 通过读取 JSON 文件
	file, err := os.Open("servers.json") // For read access.		
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Serverslice{}
	json.Unmarshal(data, &v)
	fmt.Println(v)
	fmt.Println("------------------------------")
	for _,v := range v.Servers {
		fmt.Println(v.ServerName, v.ServerIP)
	}


	// 用 map 存储
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	// 如果在我们不知道他的结构的情况下，我们把他解析到interface{}里面

	var f interface{}
	err = json.Unmarshal(b, &f)
	// 这个时候f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里

	fmt.Println(f)
	fmt.Println("------ map ------")

	f_map := f.(map[string]interface{})	// interface 转其他类型 (通过断言的方式 type assert)
	fmt.Println(f_map["Name"])
	fmt.Println(f_map["Age"])
	fmt.Println(f_map["Parents"])
}