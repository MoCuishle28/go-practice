package main

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Show() {
	for _,v := range slice {
		fmt.Println(v)
	}
}

func (s ByteSlice) change_without_point() {
	for i, _ := range s{
		s[i] += 1
	}
}

func (ps *ByteSlice) change_with_point() {
	data := *ps		// 为什么需要加这一步？ 直接 range ps 为什么出错
	for i, _ := range data{
		data[i] += 1
	}
}

func (slice ByteSlice) add(data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {  // 重新分配
		// 为了后面的增长，需分配两份。
		newSlice := make([]byte, (l+len(data))*2)
		// copy 函数是预声明的，且可用于任何切片类型。
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

func (ps *ByteSlice) add_with_point(data []byte) {
	slice := *ps

	l := len(slice)
	if l + len(data) > cap(slice) {  // 重新分配
		// 为了后面的增长，需分配两份。
		newSlice := make([]byte, (l+len(data))*2)
		// copy 函数是预声明的，且可用于任何切片类型。
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}

	*ps = slice
}

func main() {
	bs := ByteSlice {}
	bs = append(bs, 1)
	bs = append(bs, 2)
	bs = append(bs, 3)
	bs.Show()

	// 为什么 接收者是否为指针都会修改可见？
	fmt.Println("-------change_without_point----------")
	bs.change_without_point()
	bs.Show()

	fmt.Println("--------change_with_point---------")
	bs.change_with_point()
	bs.Show()

	// 添加又和需改不同
	fmt.Println("-------add_without_point-------")
	tmp := ByteSlice {0,0}
	bs.add(tmp)	//需要接收一下返回值才能成功添加
	// bs = bs.add(tmp)
	bs.Show()

	fmt.Println("-------add_with_point-------")
	bs.add_with_point(tmp)
	bs.Show()
}