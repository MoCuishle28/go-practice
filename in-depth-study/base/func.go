package main

import(
	"fmt"
	"reflect"
	"runtime"
)


func operate(a, b int, op string) (int, error) {
	var ret int
	var err error = nil
	switch op {
		case "+":
			ret = a + b
		case "-":
			ret = a - b
		case "*":
			ret = a * b
		case "/":
			ret = a / b
		default:
			err = fmt.Errorf("operate error!")
	}
	return ret, err
}


// 以函数为参数
func apply(op func(int, int) int, a, b int) int {
	// 反射拿到函数名
	p := reflect.ValueOf(op).Pointer()	// 获得函数的指针
	opName := runtime.FuncForPC(p).Name()	// 获得函数名字

	fmt.Printf("%s, Type:%T\n", opName, op)
	return op(a, b)
}


// 可变参数列表 任意多个int类型
func sum(numbers ...int) int {
	fmt.Println("Variable args list:", numbers)
	ret := 0
	for _, i := range numbers {
		fmt.Println(i)
		ret += i
	}
	return ret
}


func main() {
	res, err := operate(3, 4, "*")
	fmt.Println(res, err)

	// 传入匿名函数
	res = apply(func(a, b int) int {
		return a + b
	}, 3, 4)
	fmt.Println(res)

	fmt.Println("sum(...int):", sum(1, 2, 3, 4, 5))
}