package main

import (
	"fmt"
)

type Stringer interface{
	String() string 	// 接口包含的函数？
}

func main() {
	var value interface{}
	switch str := value.(type){
		case string:
			fmt.Println(str)
		case Stringer:
			fmt.Println(str.String())
		default:
			fmt.Println("default")
	}

	// ？？？类型断言是什么鬼？？？
	if str, ok := value.(string); ok {
		fmt.Println(str)
	} else if str, ok := value.(Stringer); ok {
		fmt.Println(str.String())
	}

	str, ok := value.(string)
	fmt.Println("end", str, ok)
}