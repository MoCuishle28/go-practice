package main

import(
	"fmt"

	"Go-practice/resk/infra/algo"
)


func main() {
	count, amount := int64(10), int64(100)*100

	arr := algo.AfterShuffle(count, amount)
	sum := int64(0)
	for _, v := range arr {
		sum += v
	}
	fmt.Println(arr)
	fmt.Println(sum, amount, sum == amount)
}