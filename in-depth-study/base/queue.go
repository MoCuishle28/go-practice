package main

import (
	"fmt"
)

/*
生成文档:
go doc 包名
godoc -http :6060
*/


// interface{}表示支持任何类型
type Queue []interface{}


func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)	// 加上取值符号才能取到slice，否则是一个slice的地址
}


func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}


func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}


// 切片是一个包含 ptr指针(指向数组某个索引) len cap 三个属性的结构体
func main() {
	q := Queue{1}		// 加不加 & 都能正常使用下面函数, 接收者加上*后 编译器似乎会自动取出地址传进函数
	q.Push(2)
	q.Push("aaa")
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}