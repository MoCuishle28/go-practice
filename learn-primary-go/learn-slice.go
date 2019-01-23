package main

import "fmt"

// 切片参数传递 本质是传递了指针指向的数组
func change_slice(s []int) {
	for i,_ := range s {
		s[i] += 1
	}
}

func main() {
	s1 := make([]int, 3, 5)
	for _,v := range s1{
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(len(s1), cap(s1), s1)

	change_slice(s1)
	fmt.Println(s1)

	fmt.Println("---测试重新分配切片---")
	s2 := s1				// 切片赋值 会引用同一个数组
	fmt.Println("s2:",s2)

	s1 = append(s1, 5)
	fmt.Println("s1:",len(s1), cap(s1), s1)

	//s1, s2 都改变了，超出容量后 不会分配新数组（cap 指的是切片对分配数组前cap个进行引用）
	// len 才是分配数组的长度
	change_slice(s1)
	fmt.Println("---超出容量后 change---")
	fmt.Println("s1:",len(s1), cap(s1), s1)
	fmt.Println("s2:",len(s2), cap(s2), s2)

	s1 = append(s1, 5)
	s1 = append(s1, 5)		// 超出长度后 分配新的数组吗？
	s1 = append(s1, 5)
	fmt.Println("s1:",len(s1), cap(s1), s1)
	fmt.Println("s2:",len(s2), cap(s2), s2)

	fmt.Println("---超出长度后 change---")
	change_slice(s1)		// 修改s1 但是s2不变 即：超出长度后，分配了新的数组给s1
	fmt.Println("s1:",len(s1), cap(s1), s1)
	fmt.Println("s2:",len(s2), cap(s2), s2)

	fmt.Println("---二维切片---")

	pic := make([][]float32, 3)		// 先make一个切片的切片
	fmt.Println(pic)
	for i := range pic {
		pic[i] = make([]float32, 3)	//再逐行分配
	}

	fmt.Println(pic)
	pic[0][0] = 5.5
	fmt.Println(pic)

	for i := range pic {
		for j := range pic[i] {
			fmt.Print(pic[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}