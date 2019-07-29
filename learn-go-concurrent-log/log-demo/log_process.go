package main

import(
	"flag"
	"strings"
	"fmt"
	"time"
	"os"
	"bufio"
	"io"
	"regexp"
	"log"
	"strconv"
	"net/url"

	"github.com/influxdata/influxdb1-client/v2"
)


// 定义接口是为了把读取与写入抽象出来，以方便支持各种读取、写入（如：文件写入、数据库写入等）
type Reader interface{
	Read(rc chan []byte)
}


type Writer interface{
	Write(wc chan *Message)
}


type LogProcess struct {
	rc chan []byte 			// 读取模块到解析模块间传递数据 (read channel)
	wc chan *Message 		// 解析模块到写入模块间解析数据 (write channel)
	read Reader 			// 读取器
	write Writer 			// 写入器
}


type ReadFromFile struct {
	path string 			// 读取文件路径
}


type WriteToInfluxDB struct {
	influxDBDsn string 		// influx data source	
}

type Message struct {
	TimeLocal	time.Time
	BytesSent	int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime float64
}


func (r *ReadFromFile) Read(rc chan []byte) {
	// 读取模块
	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)	// 移动到文件末尾
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')	// 读取一行文件内容，直到遇到换行符
		// 输入一行后要加入换行符

		//如果读取到文件末尾
		if err == io.EOF {
			time.Sleep(500*time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}


func (w *WriteToInfluxDB) Write(wc chan *Message) {
	// 写入模块

	infSli := strings.Split(w.influxDBDsn, "@")

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     infSli[0],
		Username: infSli[1],
		Password: infSli[2],
	})
	if err != nil {
		log.Fatal(err)
	}

	for v := range wc {
		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  infSli[3],	// 指定数据库
			Precision: infSli[4],	// 精度参数
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create a point and add to batch
		// Tags: Path, Method, Scheme, Status
		tags := map[string]string{"Path": v.Path, "Method": v.Method, "Scheme": v.Scheme, "Status": v.Status}
		// Fields: UpstreamTime, RequestTime, BytesSent
		fields := map[string]interface{}{
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
			"BytesSent":    v.BytesSent,
		}

		// 表名：nginx_log
		pt, err := client.NewPoint("nginx_log", tags, fields, v.TimeLocal)
		if err != nil {
			log.Fatal(err)	
		}
		bp.AddPoint(pt)		

		// 写入influxdb？
		// Write the batch
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

		log.Println("write success!")
	}
}


func (l *LogProcess) Process() {
	// 解析模块
	/**
	数据格式
	172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
	*/
	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			log.Println("FindStringSubmatch fail:", string(v))
			continue
		}

		message := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			log.Println("ParseInLocation fail:", err.Error(), ret[4])
			continue
		}
		message.TimeLocal = t

		// string 转 int
		byteSent, _ :=  strconv.Atoi(ret[8])
		message.BytesSent = byteSent

		// GET /foo?query=t HTTP/1.0
		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			log.Println("strings.Split fail", ret[6])
			continue
		}
		message.Method = reqSli[0]

		u, err := url.Parse(reqSli[1])
		if err != nil {
			log.Println("url parse fail:", err)
			continue
		}
		message.Path = u.Path

		message.Scheme = ret[5]
		message.Status = ret[7]

		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upstreamTime
		message.RequestTime = requestTime

		l.wc <- message
	}
}


func main() {
	// Grafana Default login and password admin/ admin

	// 通过命令行参数传入
	var path, influxDsn string
	flag.StringVar(&path, "path", "./access.log", "read file path")
	// 地址@用户名@密码@数据库@精度
	flag.StringVar(&influxDsn, "influxDsn", "http://127.0.0.1:8086@imooc@imoocpass@imooc@s", "influx data source")
	flag.Parse()

	r := &ReadFromFile{
		path: path,
	}

	w := &WriteToInfluxDB{
		influxDBDsn: influxDsn,
	}

	lp := &LogProcess{
		rc: make(chan []byte),
		wc: make(chan *Message),
		read: r,
		write: w,
	}

	// golang 有优化 不需要写成 (*lp).ReadFromFile() 也能正常工作
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(30*time.Second)
}