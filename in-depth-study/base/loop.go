package main


import(
	"fmt"
	"strconv"
	"os"
	"bufio"
)


func convertToBin(n int) (result string) {
	for ; n > 0; n >>= 1 {
		lsb := n & 1
		result = strconv.Itoa(lsb) + result
		fmt.Printf("Decimal:%d, Binary:%b\n", n, n)
	}
	return
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// 一行行读文件
	for scanner.Scan() {
		fmt.Println(scanner.Text())		// 读出的一行
	}
}


func main() {
	// 注意打印顺序，是先执行完调用函数，再把结果打印
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("abc.txt")
}