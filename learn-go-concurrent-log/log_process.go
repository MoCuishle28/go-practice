package main

import(
	"strings"
	"fmt"
	"time"
)


type LogProcess struct {
	rc chan string 			// 读取模块到解析模块间传递数据 (read channel)
	wc chan string 			// 解析模块到写入模块间解析数据 (write channel)
	path string 			// 读取文件路径
	influxDBDsn string 		// influx data source
}


func (l *LogProcess) ReadFormFile() {
	// 读取模块
	line := "message"
	l.rc <- line
}


func (l *LogProcess) Process() {
	// 解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}


func (l *LogProcess) WriteToInfluxDB() {
	// 写入模块
	fmt.Println(<-l.wc)
}


func main() {
	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
		path: "/access.log",
		influxDBDsn: "username&password...",
	}

	// golang 有优化 不需要写成 (*lp).ReadFormFile() 也能正常工作
	go lp.ReadFormFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1*time.Second)
}