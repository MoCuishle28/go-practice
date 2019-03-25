package main
import (
	"bufio"
	"fmt"
	"os"
)


// 传递的是引用的副本
func changeMap(counts map[string]int) {
	for k, _ := range counts {
		counts[k]+= 100
	}
}


func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {			// 读入下一行，并移除行末的换行符
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

	changeMap(counts)

	fmt.Println(counts)

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}