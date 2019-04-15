package main


import(
	"fmt"
	"time"
)


func main() {

	signal := make(chan int)
	closeSignal := make(chan int)

	go func(signal chan int) {
		for {
			<-signal
			fmt.Println("eat")
		}
	} (signal)


	go func(signal chan int, colseSignal chan int) {
		for i := 0 ; i < 5 ; i++{
			time.Sleep(2*time.Second)
			fmt.Println("put ", i)
			signal <- 1
		}
		time.Sleep(time.Second)
		closeSignal <- 1
	} (signal, closeSignal)

	exit := false
	for{
		if exit {
			break
		}
		select {
		case <-closeSignal:
			fmt.Println("close")
			exit = true
		}
	}
}