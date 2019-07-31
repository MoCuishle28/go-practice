package main

import(
	"fmt"
	"math/rand"
	"time"
)


// 返回的channel只能取数据  只能发数据的channel写作 chan<- int(一个在chan左边表示出 一个在chan右边表示进)
func generator() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 随机随眠1500毫秒内
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}


func main() {
	var c1, c2 <-chan int 		// c1 and c2 == nil
	c1, c2 = generator(), generator()
	exit := time.After(10 * time.Second)	// 10s 后输入一个信号给 exit 管道
	tick := time.Tick(time.Second)			// 计时器 每秒一个信号
	for {
		select {
			// 一般会将生产出来的数据存在某个队列里 等待消耗
			case n := <-c1:
				fmt.Println("Received from c1: ", n)
			case n := <-c2:
				fmt.Println("Received from c2: ", n)
			// case <-tick:
			// 	fmt.Println("tick tock!")
			case <-time.After(800*time.Millisecond):	// 800毫秒没有收到数据认为超时 (每次select开始计时)
				// 上面的10s是整个运行时间
				fmt.Println("timeout!")
			case <-exit:
				fmt.Println("bye bye!")
				return
			// default:	// 都没收到数据
			// 	fmt.Println("No value eceived...")
		}
	}
}