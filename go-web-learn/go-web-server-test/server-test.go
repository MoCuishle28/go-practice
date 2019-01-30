package main

/**
	net/http包 搭建 go web 服务
*/

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"log"
	"os"
	"time"
	"io"
	"strconv"
	"crypto/md5" 
)


func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()			// 解析url传递的参数，对于POST则解析响应包的主体 (默认不解析)
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println("start sayHelloName func...")
	fmt.Println(r.Form)		// 是一个 map 吧
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form['url_long'] =",r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, " "), " | ", v)
	}
	fmt.Fprintf(w, "Hello astaxie!")		// 写入到 w 会返回输出到客户端
	fmt.Println("Hello End")
}


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start login...")		
	fmt.Println("method:", r.Method) 		// 获取请求的方法
	r.ParseForm()							// 不设置解析是不会自动解析的
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")	// 解析模板？
		log.Println(t.Execute(w, nil))				// log？
	} else {
		fmt.Println("username:", r.Form.Get("username"))	// 用Get 为空时 会返回空值
		fmt.Println("password:", r.Form.Get("password"))	
		fmt.Println("不存在表单项", r.Form.Get("nilItem"))
		fmt.Println("===")
		for k,v := range r.Form {	// 每一项 v 都是切片
			fmt.Println(k, v)
		}
		fmt.Println("===")
	}
	fmt.Println("login End...")
}


func main() {
	http.HandleFunc("/", sayHelloName)			// 设置访问路由
	http.HandleFunc("/login", login)         	// 设置访问路由
	http.HandleFunc("/upload", upload)			// 上传文件
	// 第二个参数是 handler(nil 则默认为DefaultServeMux) 什么意思？ 是用来设置路由的
	err := http.ListenAndServe(":9090", nil)	// 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		// 上传的文件存储在maxMemory大小的内存里面
		// 如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中。(参数为 maxMemory )
		r.ParseMultipartForm(32 << 20)	// 把上传的文件存储在内存和临时文件中
		file, handler, err := r.FormFile("uploadfile")		// 获取文件句柄，然后对文件进行存储等处理
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		// 此处假设当前目录下已存在test目录
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}