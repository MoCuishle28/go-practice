package main

import "fmt"

func change_slice(slice []int) {
	slice = append(slice, 100)
	for i,v := range slice {
		v += 20
		slice[i] = v
	}
	fmt.Println("in change_slice:",slice)
}

func main() {
	var p *[]int = new([]int)       // 分配切片结构； *p == nil；基本没用
	var v  []int = make([]int, 10) // 切片 v 现在引用了一个具有 10 个 int 元素的新数组

	fmt.Println(p, v)

	// 没必要的复杂：
	// var p *[]int = new([]int)
	// *p = make([]int, 100, 100)

	// 为什么？？？？？？？？？？
	/*
	切片是对数组的一段数据的引用
	在 main 中的 slice 引用了 一个长度为10的数组的前5个元素
	在 change_slice 中 slice 与 maim 中的是引用同一段数组 因此可以修改值
	但是在 change_slice 中 append 的值在同一段数组的第6个元素的位置
	而 main 中的 slice 只能引用到前5个 所以在main中打印slice还是只能看到修改后的前5个元素
	*/
	slice := make([]int, 5, 10)
	change_slice(slice)
	fmt.Println(slice)

	fmt.Println("---")
	arr := make([]int, 2, 3)
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 1)
	arr = append(arr, 1)
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 1)
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 1)
	arr = append(arr, 1)
	fmt.Println(arr, len(arr), cap(arr))
}