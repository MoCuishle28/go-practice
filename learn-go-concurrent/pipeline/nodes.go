package pipeline
// 做一个库

import (
	"sort"
	"io"
	"math/rand"
	"encoding/binary"
)


// a是可变长参数  返回值是一个只能出的chan
func ArraySource(a ... int) <-chan int {
	out := make(chan int)
	go func() {
		for _,v := range a {
			out <- v
		}
		close(out)		// 有明显的结束 需要close
	}()					// 用匿名函数创建一个 goroutine
	return out
}

// in是只能输入的chan
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	// 开了一个goroutine 后不会等待 会直接返回out
	go func () {
		// read into menory
		a := [] int {}
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a)
		// output
		for _,v := range a {
			out <- v
		}
		close(out)
	} ()
	return out
}

func Merge(in1, in2 <- chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// 此处会等待chan中有数据才开始读数据
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		// 不关闭则导致main中的range会不知道什么时候停止
		close(out)
	} ()
	return out
}

// 读数据 chunSize 规定能读的块大小
func ReaderSource(reader io.Reader, chunSize int) <-chan int {
	// 加上缓冲区(1024kb) 可以不用发一个 就要求接收一个 因此能减少切换的频率，加快速度
	out := make(chan int, 1024)
	go func () {
		// 开一个大小为8的
		buffer := make([] byte, 8)
		bytesRead := 0
		for {
			// n返回多少个字节 err表示是否有错误(读到EOF就是有错\未读满8字节)
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunSize != -1 && bytesRead >= chunSize){
				break
			}
		}
		close(out)
	} ()
	return out
}

// 写数据到sink节点
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	} ()
	return out
}

// 多路两两归并
func MergeN(inputs ... <-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	// merge inputs[0..m) and inputs[m..end)
	// 加 ... 是为了让参数在语法上是同一类型
	return Merge( 
			MergeN(inputs[:m]...), 
			MergeN(inputs[m:]...))
}