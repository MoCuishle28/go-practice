package main

import (
	"fmt"
	"os"
	"Go-practice/learn-go-concurrent/pipeline"
	"bufio"
	"strconv"
)

func main() {
	// 大小512的文件 分成4块进行外部排序

	// 本地版
	p := createPipeline("small.in", 512, 4)

	// 网络版
	// p := createNetworlPipeline("small.in", 512, 4)

	writeToFile(p, "small.out")
	printFile("small.out")
}

func createPipeline(filename string, fileSize, chunCount int) <-chan int {
	/*
	chunCount: 分为几块读取
	*/
	chunSize := fileSize / chunCount
	sortResults := [] <-chan int{}	// 存放每一块数据的chan

	for i := 0; i < chunCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		// 定位到要读的那一块的开头 Go对类型检查很严格（第一个参数一定要int64）
		file.Seek(int64(i * chunSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunSize)
		// 每一块单独进行InMemSort
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	// 单独排序完了 再进行N路的两两归并
	return pipeline.MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v := range p{
		fmt.Println(v)
	}
}

func createNetworlPipeline(filename string, fileSize, chunCount int) <-chan int {
	chunSize := fileSize / chunCount
	sortAddr := [] string {}

	for i := 0; i < chunCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		// 定位到要读的那一块的开头 Go对类型检查很严格（第一个参数一定要int64）
		file.Seek(int64(i * chunSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunSize)
		
		addr := ":" + strconv.Itoa(7000 + i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	sortResults := [] <-chan int{}
	for _, addr := range sortAddr{
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}