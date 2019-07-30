package main

import(
	"fmt"
)


func Append(s []int, x int) {
	s = append(s, x)
	fmt.Println("Append:", s)
}

func Append_ref(s *[]int, x int) {
	*s = append(*s, x)
	fmt.Println("Append_ref:", s)
}


func main() {
	arr := [7]int{0, 1, 2, 3, 4, 5, 6}
	s := arr[0:6]
	fmt.Println(s, arr)

	Append(s, 90)
	fmt.Println(s, arr)
	Append(s, 91)
	fmt.Println(s, arr)
	Append(s, 92)
	fmt.Println(s, arr)

	fmt.Println("---ref---")
	arr[6] = 6	
	fmt.Println(s, arr)
	Append_ref(&s, 90)
	fmt.Println(s, arr)
	Append_ref(&s, 91)
	fmt.Println(s, arr)
	Append_ref(&s, 92)
	fmt.Println(s, arr)

	s[0] = 100
	fmt.Println(s, arr)		// 此时s已经不再引用arr数组了
}