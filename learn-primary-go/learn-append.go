package main

import "fmt"

func main() {
	a := [] int {1,2,3}
	b := a
	a = append(a, 4, 5, 6)
	fmt.Println("a:",a)
	fmt.Println("b:",b)

	fmt.Println("---------------")
	a = b
	a = append(a, b...)
	fmt.Println("a:",a)

	for i,_ := range a{
		a[i] = (i+1)*10
	}

	fmt.Println("a:",a)
	fmt.Println("b:",b)

	// 用append进行元素删除
	fmt.Println("---Delete---")
	fmt.Println(a)
	a = append(a[:3], a[4:]...)
	fmt.Println(a)
}