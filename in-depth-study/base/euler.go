package main

import(
	"fmt"
	"math/cmplx"
	"math"
	"io/ioutil"
)


func triangle() {
	// 类型转换 只能强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4	// 若常量不指定类型，那既可以是int也可是float（只是文本）
	var c int
	fmt.Printf("type:%T, %T, %T\n", a, b, filename)
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
	fmt.Printf("type:%T, %T, %T\n", a, b, filename)
}


func readFile() {
	// 判断语句
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
	}
	// 这种写法contents变量只能在if语句(包括关联的else、 else if)里使用
	// fmt.Println(contents) // 会报错
}


func enums() {
	// 枚举变量
	const(
		// 还可以用iota 表示这组常量是自增值
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// iota复杂用法 后面一组会用同一个公式 其中iota自增
	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 验证欧拉公式 e^i*Pi + 1 = 0 (1i 表示虚数 i)
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	// 还可以用Exp 以E为底数的x次方计算 得到结果很接近0
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)

	triangle()
	consts()
	enums()
	readFile()
}