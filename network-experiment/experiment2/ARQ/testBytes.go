package main


import(
	"fmt"
	"strings"
)


func main() {
	s := "abcdefg"
	
	byteArr := []byte(s)

	for _, v := range byteArr {
		fmt.Println(v)
	}
	fmt.Println(byteArr[0])
}