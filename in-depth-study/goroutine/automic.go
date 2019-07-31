package main

import(
	"sync"
	"fmt"
	"time"
)

/*
传统同步机制例子
*/


// 模拟实现一个原子加
type atomicInt struct {
	value int
	lock sync.Mutex
}


func (a *atomicInt) increment() {
	fmt.Println("safe increment!")
	// 加上匿名函数并在函数体内用锁 可以使得整块函数体执行被锁住(安全)
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	} ()
}


func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}


func main() {
	var a atomicInt
	a.value = 1
	a.increment()
	fmt.Println(a.get())
	time.Sleep(time.Second)
}