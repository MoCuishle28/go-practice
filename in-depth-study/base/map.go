package main


import(
	"fmt"
)


func setMap(m map[string]string) {
	m["name"] = "bbbbbbbbbbbbbb"
}


func main() {
	/*
	作为key的类型必须能比较相等
	除了slice、map、function的内建类型都能作为key
	struct类型只要不包含上述三种也可以作为key
	*/
	m := map[string]string {
		"name":"a",
		"course":"golang",
	}
	fmt.Println(m, m["null"] == "")

	// value为一个map的map
	m1 := map[string] map[string]string{
		"a":m,
	}
	fmt.Println(m1)

	// 初始化不同，初值代表含义不同
	m2 := map[int]int{}		// m2 != nil  m2 是empty map
	var m3 map[int]int 		// m3 == nil
	m4 := make(map[int]int)	// m4 != nil  m4 是empty map
	fmt.Println(m2, m2 == nil, m3, m3 == nil, m4, m4 == nil)

	// 传进去的是地址, map类型应该是一个带有指针指向真正哈希表的结构体，所以传递时把指针值赋值过去了
	setMap(m)
	fmt.Println(m)
}