package main

import(
	"net/http"
	"strings"
	"crypto/md5"
	"time"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"encoding/json"
)


func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}


func main() {
	http.HandleFunc("/sayHello", sayHello)

	// 实现读取文件的handler
	fileHandler := http.FileServer(http.Dir("video"))	// 读取文件夹video下的文件
	// 注册handler
	http.Handle("/video/", http.StripPrefix("/video/", fileHandler))	// 除去video前缀

	// 注册上传文件
	http.HandleFunc("/api/upload", uploadHandler)

	// 注册获取文件列表
	http.HandleFunc("/api/list", getFileListHandler)

	http.ListenAndServe(":8090", nil)
}


// 上传视频文件
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 限制客户端上次视频大小 10MB
	// 截取前10MB内容，若可以按照http的body体进行解析，则认为是小于等于10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	err := r.ParseMultipartForm(10*1024*1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// 从请求体中获取上传文件
	file, fileHandler, err := r.FormFile("uploadFile")	// key为uploadFile

	// 检查文件是否为MP4类型
	ret := strings.HasSuffix(fileHandler.Filename, ".mp4")
	if ret == false {
		http.Error(w, "not mp4", http.StatusInternalServerError)
		return 
	}

	// 获取随机名称，给文件重命名
	md5Byte := md5.Sum([]byte(fileHandler.Filename + time.Now().String()))
	md5Str := fmt.Sprintf("%x", md5Byte)
	newFileName := md5Str + ".mp4"

	// 写入文件到磁盘video目录下
	dst, err := os.Create("./video/"+newFileName)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 	
	}
	defer file.Close()
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	return
}


// 获取视频文件列表
func getFileListHandler(w http.ResponseWriter, r *http.Request) {
	// 设置请求的域名可以为任意域名。 TODO？？？
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 获取指定文件夹下所有文件, 返回一个数组，其中包含每个文件的路径
	files, _ := filepath.Glob("video/*")
	var ret []string
	for _, file := range files {
		// 将每个文件路径专程接口指定格式的url
		// Host是地址和端口号 filepath.Base(file)是文件名
		ret = append(ret, "http://" + r.Host+ "/video/" + filepath.Base(file))
	}
	// 转成接口指定的json格式
	retJson, _ := json.Marshal(ret)
	w.Write(retJson)
	return
}