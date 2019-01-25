package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

func to_hash(s string) string {
	// 需要参数为字节切片 返回的是32个元素的字节数组
	hashInBytes := sha256.Sum256([] byte(s))
	// 字节切片转换为字符串 hashInBytes是字节数组 需要先转换为字节切片
	retString := hex.EncodeToString(hashInBytes[:])
	fmt.Println(hashInBytes)
	return retString
}

func main() {
	fmt.Println(to_hash("test1"))
}