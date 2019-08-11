package main

import(
	"fmt"

	"Go-practice/resk/infra/algo"
)


func main() {
	count, amount := int64(10), int64(100)*100
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := algo.DoubleAverage(count-i, amount)
		amount = amount - x
		sum += x
		fmt.Printf("%d, ", x)
	}
	fmt.Println()
	fmt.Println(remain, sum, remain == sum)
}