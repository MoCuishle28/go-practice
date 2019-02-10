package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')	 	// 为什么要用单引号？
	fmt.Println(string(str))
	
	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	// g 表示格式 12是精度 64指 float64
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(-1234, 10)
	d := strconv.FormatUint(12345, 10)
	// Itoa是FormatInt(i, 10) 的简写。
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// Parse 系列函数把字符串转换为其他类型
	a0, err := strconv.ParseBool("false")
	checkError(err)

	b0, err := strconv.ParseFloat("123.23", 64)
	checkError(err)

	c0, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)

	d0, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)

	e0, err := strconv.Atoi("1023")
	checkError(err)

	fmt.Println(a0, b0, c0, d0, e0)
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}