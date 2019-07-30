package main

import(
	"fmt"
)


// 不加 [长度] 就是slice
func updateSlice(s []int) {
	// slice结构体内部有一个指向底层数组某个索引位置的指针, slice修改值根据该指针修改
	// 所以在函数内修改，底层数据也会跟着修改
	s[0] = 100
}


func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}	// 让编译器数
	fmt.Println(arr1, arr2)	
	fmt.Println(arr3, len(arr3))

	var grid [4][5]int 		// 二维 4个长度为5的数组
	fmt.Println(grid)

	fmt.Println("------------")
	// Go语言指针与数组在函数中的应用 -> pointer.go
	// Go一般不直接使用数组，用slice
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]		// 切片
	fmt.Println(s)
	// slice 不是值类型，而是一个视图
	updateSlice(s)
	fmt.Println("updateSlice(s):", s, arr)

	fmt.Println("------------")
	// slice虽然不能直接用[i]取出超过索引的值，但是能通过 re slice 取到
	// slice 可以向后扩展，看到slice len之后的数，但最长不超过cap（不能向前扩展）
	arr[2], arr[2] = 0, 2
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2, arr)

	// append之后 底层的数组也变了
	fmt.Println("------------")
	fmt.Println(arr)
	s3 := append(s2, 10)
	fmt.Println(s2[:3], s3, arr)

	fmt.Println("------------")
	// 但是，当append长度超过原数组(超过cap)，则会新开一个数组作为slice的引用，并将原先的元素拷贝过去
	// 每次装不下会创建cap翻倍的数组被新slice引用
	s4 := append(s3, 11)
	fmt.Println(s3[:3], s4, arr)	// s3不能取到[:4]了，因为此时s3与s4引用的数组不是同一个


	// slice的删除
	fmt.Println("-------------slice 删除操作----------")
	s = make([]int, 3, 5)
	for i, _ := range s {
		s[i] = i
	}
	fmt.Println(s)
	s = append(s, 3)
	s = append(s, 4)
	fmt.Println(s)
	s1 = append(s[:2], s[3:]...)	// slice删除索引2的值 后一个是可变长度参数
	fmt.Println(s1, s)				//s引用的数组和s1引用的是同一个 也跟着变了(相当于向前复制了一位，但最后一位不变)
		
}