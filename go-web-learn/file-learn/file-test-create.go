package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)	// 返回 *File

	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test! by WriteString\r\n")		// 写入文件
		fout.Write([]byte("Just a test! by Write\r\n"))		// 写入文件
	}
}