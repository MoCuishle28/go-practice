package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	// 例子用了ini格式配置
	conf, err := config.NewConfig("ini", "./logcollect.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	
	listen_ip := conf.String("server::listen_ip")
	fmt.Println("IP:", listen_ip)

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}
	fmt.Println("Port:", port)

	log_level := conf.String("log::log_level")	// string 不返回 err
	fmt.Println("log_level:", log_level)

	log_path := conf.String("log::log_path")
	fmt.Println("log_path:", log_path)
}
