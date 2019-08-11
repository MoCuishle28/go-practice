package main

import(
	"fmt"

	"Go-practice/resk/infra/algo"
)


// 测试简单随机算法 可以看出明显后面的金额较少
func main() {
	// 100个红包 总额100元
	count, amount := int64(12), int64(128)*100

	remain := amount
	var sum int64 = 0
	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count-i, amount)
		sum += x
		amount = amount - x
		fmt.Print(float64(x)/float64(100), ", ")
	}
	fmt.Println()
	fmt.Println(sum, remain, sum == remain)
}