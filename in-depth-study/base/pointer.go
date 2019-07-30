package main

import(
	"fmt"
)


func f0(a *int) {
	*a++
}


func f1(arr [3]int) {
	// 也只是接受值的复制
	for i := range arr {
		arr[i]++
	}
	fmt.Println("f1:", arr)
}


// 指针 指向数组的指针
func f2(arr *[3]int) {
	for i := range arr {
		arr[i]++
	}
	fmt.Println("f2:", arr)	
}


func main() {
	// Golang只有值传递
	a := 1
	f0(&a)
	fmt.Println(a)

	arr1 := [3]int{1, 2, 3}
	f1(arr1)
	fmt.Println(arr1)

	f2(&arr1)	//复制数组的地址过去
	fmt.Println(arr1)
}