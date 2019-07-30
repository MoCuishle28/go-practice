package main

import(
	// "errors"
	"fmt"
)


func tryRecover() {
	// defer中执行recover
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	} ()

	// panic如果不遇到recover就会直接退出程序
	// panic(errors.New("this is an error"))

	// b := 0
	// a := 5 / b
	// fmt.Println(a)

	// panic一个不是错误的东西，defer 执行的recover就会进入else 再次panic
	panic(123)
}


func main() {
	tryRecover()
}