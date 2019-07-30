package main

import(
	"fmt"
	"unicode/utf8"
)


func main() {
	s := "哈哈哈hhh"		// utf-8 编码, 每个中文3字节
	fmt.Println(len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		// i的索引也会跟着跳
		fmt.Printf("(%d,%X,%T) ", i, ch, ch)	// ch 是一个rune类型 (4字节整数 int32)
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))	// 获得字符数量
	ch, size := utf8.DecodeRune([]byte(s))
	fmt.Println(ch, size)
	fmt.Printf("%c, %T\n", ch, ch)

	for i, ch := range []rune(s) {			// 这样i就可以从0~len
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Println(i, string(ch))
	}
	fmt.Println()
}