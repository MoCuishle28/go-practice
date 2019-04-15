package main

import(
	"fmt"
    "time"
)


func main() {
	t1 := time.NewTimer(5*time.Second)
	// t2 := time.NewTimer(3*time.Second)

	for {
		// 先执行这里 然后阻塞在 select 中等待某个 case 的执行
		// (if 没有default, else 有default就执行default)
		fmt.Println(t1.C)
		
		select {
	
		case <-t1.C:
			fmt.Println("5s")
			t1.Reset(time.Second*5)

		// default:
		// 	fmt.Println("default")
	
		// case <-t2.C:
		// 	fmt.Println("3s")
		// 	t2.Reset(time.Second*3)
	
		}
	}

	fmt.Println("END")
}