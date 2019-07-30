package main

import (
	"fmt"
)


// 实现接口的结构体
type MockRetriever struct {
	Contents string
}


// 只要实现了Get方法(Retriever内唯一的方法)，就认为实现了 Retriever 接口
func (r MockRetriever) Get(url string) string {
	return r.Contents
}


type Retriever interface {
	// interface内不需要加func关键字，因为interfa里面全是函数
	Get(url string) string
	// 接口里面也可以放接口，直接写接口名字就行
}


func download(r Retriever) string {
	return r.Get("www.imooc.com")
}


func inspect(r Retriever) {
	// 看看接口内部存了 类型 和 对象拷贝（也可以是对象指针，由创建时决定）
	fmt.Printf("inspect!!! type:%T, value:%v\n", r, r)
	// 看到接口属于指针还是值
	// type switch
	switch v := r.(type) {
		case MockRetriever:
			fmt.Println("contents:", v.Contents)
		case *MockRetriever:
			fmt.Println("* contents:", v.Contents)
	}
}


func main() {
	var r Retriever
	r = MockRetriever{"this is duck"}
	fmt.Println(download(r))

	// 因为 MockRetriever 实现了函数参数所要求的接口，所以可以直接当作duck传入
	fmt.Println(
		download(MockRetriever{"this is another duck"}))

	inspect(r)

	r = &MockRetriever{"interface ref!"}
	inspect(r)

	// 另一种检查类型  (type assertion)
	rr, ok := r.(Retriever)
	fmt.Println(rr, ok)

	// interface{}表示任何类型
	var a interface{}
	a = 1					// 赋值后成为确定类型?
	fmt.Printf("%T,%d\n", a, a)
	// 可以强转为确定类型
	b := a.(int)
	fmt.Printf("%T,%d\n", b, b)
}