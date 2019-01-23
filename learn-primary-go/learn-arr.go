package main

import "fmt"

// 可以预设返回变量名 这样可以写一个空的 return 作为返回
func Sum(a *[3]float64) (sum float64) {
	for i, v := range *a {
		sum += v
		fmt.Println(a[i])
		a[i] += 1
	}
	return
}

func change_arr(a *[3]int) {
	for i,_ := range *a {
		a[i] += 1
	}
}

func main() {
	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array)  // 注意显式的取址操作
	fmt.Println(x)
	fmt.Println(array)

	fmt.Println("---")
	a := [3]int{5,6,7}
	fmt.Println("raw a：", a)
	change_arr(&a)
	fmt.Println("after change a：", a)
}