package main

import (
	_ "fmt"
	"log"

	// 服务器端编程
	"html/template"
	"net/http"

	// util
	"sync"

	// service 包
	"Go-practice/orderManager/service"
)

// 用于渲染登录页模板的信息
type LoginMsg struct {
	Msg string
}

// root 用户登录标志
type Root struct {
	username string  // 未登录时为 "null" 登录后为用户名
	lock sync.Mutex
}

// header
type Header struct {
	Username string
	Index string
}


var root *Root 			// root 用户登录标志


func init() {
	root = &Root{username:"null"}
}


// 主函数
func main() {
	// 不启动静态文件服务就无法加载 css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务

	http.HandleFunc("/login", login)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)	// 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func index(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}
	t, _ := template.ParseFiles("templates/index.html", "templates/header.html")
	header := Header{Username:root.username, Index:"0"}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "index", nil)
}


func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {	// POST
		username := r.Form.Get("username")
		pwd := r.Form.Get("pwd")

		msg := LoginMsg{}
		ret := service.ValidLogin(username, pwd)
		switch(ret){
			case -1:
				msg.Msg = "其他错误!"
			case 1:
				root.lock.Lock()
				root.username = username
    			root.lock.Unlock()
    			index(w, r)
    			return
			case 2:
				msg.Msg = "用户名错误!"
			case 3:
				msg.Msg = "密码错误!"
			default:
				msg.Msg = "异常!"
		}
    	t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, msg)
	}
}