package main

import "fmt"

type Book struct{
	id string
	price float64
}

// map是引用类型
func change(d map[string] int) {
	d["a"] += 20
}

func main() {
	d1 := map [string] int {"a":1, "b":2 }
	fmt.Println(d1)
	change(d1)
	fmt.Println(d1)

	// 奇怪的键类型... （数组也支持） 这两个都是值引用,所以值相等就能在map中取到
	book := new(Book)
	book.id = "12306"
	book.price = 120.00
	d2 := map [Book] string { *book:"a"}
	fmt.Println(d2)

	tmp_book := Book{ id:"12306", price:120.0 }
	fmt.Println("d2[tmp_book] =",d2[tmp_book])
	d2[tmp_book] = "b"
	fmt.Println(d2)

	fmt.Println("---array as key---")
	// 必须指定长度 不然就成了切片了
	arr1 := [3] int {1,2,3}
	d3 := map [[3]int] string {}
	d3[arr1] = "aaa"
	arr2 := [3] int {1,2,3}
	fmt.Println(d3)
	fmt.Println("d3[arr2] =", d3[arr2])

	fmt.Println("---若不存在key 会返回零值---")
	// 若不存在key 会返回零值
	d := map [string] bool {"mocuishle":true}
	fmt.Println("d[Jone] =",d["jone"], " but d['mocuishle'] =", d["mocuishle"])
}