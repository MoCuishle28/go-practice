package main

import (
	"Go-practice/learn-go-concurrent/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	// MergeDemo()	//不是在文件里读取的demo

	const filename = "small.in"
	const n = 64

	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()	//类似于java的finally
	
	p := pipeline.RandomSource(n)

	// bufio.NewWriter(file) 包装一下 给一个缓冲区加快速度
	writer := bufio.NewWriter(file)
	pipeline.WriterSink( writer, p)
	writer.Flush()	// 有缓冲区后 要Flush 确保所有数据都写到外存

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count == 100{
			break
		}
	}
}


func MergeDemo() {
	// p is channel
	p := pipeline.Merge(
			pipeline.InMemSort(
				pipeline.ArraySource(3, 2, 6, 7, 4)),
			pipeline.InMemSort(
				pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))

	// for {
	// 	// 若 p 这个chan 被close了 则ok为false
	// 	if num, ok := <-p ; ok {
	// 		fmt.Println(num)
	// 	} else{
	// 		break
	// 	}
	// }

	// 也可以用range （这种方式 发送方一定要close 否则range不知道什么时候结束）
	// range 会等待未准备好的chan数据
	for v := range p {
		fmt.Println(v)
	}
}