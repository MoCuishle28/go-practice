package main

import (
	"fmt"
	"time"
	"sync"
)

/*
Go的协程是基于CSP模型设计出来的

对于不带缓冲区的channel 只发送数据进入channel 而没有将channel内数据发送出去 会导致死锁
	即:channel内有数据就必须被拿出来使用

若无缓冲区 每次输入数据到channel 都导致协程切换(切换到从channel中读取数据那里) 效率较低

channel 可以close 告诉接收方明确的结束信息
close之后接收方会一直收到零值, 即:数值类型会是0 字符串类型会是空串 等等
	可以通过	 n, ok := <-c 		// channel关闭后 ok会收到false
	判断channel是否关闭
*/


type worker struct {
	in chan int
	// done chan bool

	// wg *sync.WaitGroup 	// 可以使用库实现的解决方案
	done func()		// 进一步包装一下 WaitGroup 用一个函数存 WaitGroup 的 Done()
}


// func doWork(id int, in chan int, done chan bool)
func doWork(id int, w worker) {
	// 一直接收 直到channel关闭
	for n := range w.in {
		fmt.Printf("id:%d, received:%c\n", id, n)

		// 因为输入数据后会阻塞等待数据被取出
		// 所以要另开一个协程发送完成通知 否则会一直等着done被取出 而大写字符发送进in也等待被取出 导致死锁
		// go func() {
		// 	done <- true 		// 每打印完一个就通知一下
		// } ()

		// 除了上面done的方法 还可以用库提供的 WaitGroup
		w.done()
	}
}


// 必需有消费channel内数据的逻辑 不然编译器认为会一直等待channel被消费而导致死锁
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// done: make(chan bool),
		// wg: wg,
		done: func() {
			wg.Done()
		},
	}
	// go doWork(id, w.in, w.done)
	go doWork(id, w)
	return w
}


func chanDemo() {
	var workers [10] worker
	var wg  sync.WaitGroup		// 引用一个系统提供的 WaitGroup ?

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)	 	// 小写和大写字母 一共20个任务  也可以在下面每次循环就 Add(1)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()

	// // 全部发送完再接收打印完毕信息
	// for _, workers := range workers {
	// 	// 每个in都输入了两个 所以接收两个done
	// 	<- workers.done
	// 	<- workers.done
	// }
}


// 带缓冲区的channel
func bufferedChannel() {
	// 有缓冲区之后 不必每次输入channel都必须有读取
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	// c <- 4 	//缓冲区大小3 再发4会死锁

	// 关闭channel 之后会发送零值
	close(c)

	for i := 0; i < 10; i++ {
		n, ok := <-c 		// channel关闭后 ok会收到false
		fmt.Println(n, ok)
	}
}


func main() {
	chanDemo()

	bufferedChannel()

	time.Sleep(time.Millisecond)
}