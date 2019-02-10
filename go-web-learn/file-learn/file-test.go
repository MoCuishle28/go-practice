package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("astaxie0", 0777)					// 0777 是权限设置 (perm)
	os.MkdirAll("astaxie1/test1/test2", 0777)	// 多级目录
	err := os.Remove("astaxie0")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie1")			// 删除多级目录
}