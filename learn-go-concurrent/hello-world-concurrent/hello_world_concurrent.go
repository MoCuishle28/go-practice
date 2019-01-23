package main

import (
	"fmt"
	"time"
)

func main() {
	// 并发版本的hellow world
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		// go start a goroutine 类似于开了个线程(但并不是线程)
		go printHelloWorld(i, ch)
	}

	for{
		// 不断读取 channel 的数据
		msg := <- ch
		fmt.Println(msg)
	}

	// (没有无限循环的前提下)若不睡眠 则主线程会比	goroutine 早结束 即什么都没来得及输出
	time.Sleep(100*time.Millisecond)
}

func printHelloWorld(index int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello %d goroutine!\n", index)
	}
}