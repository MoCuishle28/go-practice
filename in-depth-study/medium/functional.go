package main


import(
	"fmt"
	"io"
	"strings"
	"bufio"
)


// 返回值是一个函数
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v 	//sum 是被引用的外部自由变量
		return sum
	}
}


func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}


// 让斐波那契额数列函数能打印，只要实现打印接口
// 定义类型 intGen 是一种函数
type intGen func() int

// 实现接口就能打印，任何变量都能实现接口(包括函数)
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()		// 取得下一个元素
	// 设定结束点，不然斐波那契额数列读不完
	if next > 1000 {
		return 0, io.EOF
	}
	// 然后要写进p这个字节数组, 返回写入几个字节n

	// 先转成字符串，再让strings.NewReader(s).Read(p)来代理实现Read
	// 不然手动写入p太麻烦了
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}


func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {	// 一直读取到末尾 EOF
		fmt.Println(scanner.Text())
	}
}

func main() {
	a := adder()
	// 最终算出0~9的累加和
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d\n", i, a(i))
	}

	f := fibonacci()
	printFileContents(f)	// f已经实现了Reader接口
	// fmt.Println(f())	// 1
	// fmt.Println(f())	// 1
	// fmt.Println(f())	// 2
	// fmt.Println(f())	// 3
	// fmt.Println(f())	// 5
	// fmt.Println(f())	// 8
	// fmt.Println(f()) 	// 13
	// fmt.Println(f())	// 21
}