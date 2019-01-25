package main

import (
	"fmt"
	"time"
)

func do_something(c chan int) {
	go func () {
		fmt.Println("in do_something Sleep 2 Second...")
		time.Sleep(2 * time.Second)
		c <- 1
		fmt.Println("do_something end")
	} ()
}

func recive_chan(c chan int) {
	go func () {
		fmt.Println("in recive_chan Sleep 3 Second...")
		time.Sleep(3 * time.Second)	
		fmt.Println("recive chan :",<-c)	// 不带缓冲区，则会一直阻塞 等待信道有数据
		fmt.Println("end recive_chan")
	} ()
}

func main() {
	// c := make(chan int)
	c := make(chan int, 10) //若加上缓冲区, 则发送者只要把数据放入信道就可以结束阻塞（缓冲区未满的情况下）
	do_something(c)
	fmt.Println("continue")
	recive_chan(c)
	time.Sleep(10 * time.Second)	// 若主线程结束了， 则go程也会销毁
	fmt.Println("main end")
}