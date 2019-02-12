package main

import (
	"log"

	// 服务器端编程
	"html/template"
	"net/http"

	// util
	"sync"

	// service 包
	"Go-practice/orderManager/service"
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
)


// root 用户登录标志
type Root struct {
	username string  // 未登录时为 "null" 登录后为用户名
	lock sync.Mutex
}

// 以下为模板渲染内容
// 用于渲染登录页模板的信息
type LoginMsg struct {
	Msg string
}

// header
type Header struct {
	Username string
	Index string
}

type Index struct {
	Orders_list *[]entity.Orders
}

type DetialOrder struct {
	Detial_order_list *[]entity.Detial_order
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
	http.HandleFunc("/ordersdetial", ordersDetial)
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

	orders_list := dao.QueryOrders()

	t, _ := template.ParseFiles("templates/index.html", "templates/header.html")
	header := Header{Username:root.username, Index:"0"}
	index := Index{Orders_list:orders_list}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "index", index)
}


func ordersDetial(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	r.ParseForm()
	oid := r.Form.Get("oid")
	detial_order_list := dao.QueryOrdersDetial(oid)

	t, _ := template.ParseFiles("templates/ordersDetial.html", "templates/header.html")
	header := Header{Username:root.username, Index:"0"}
	detial_order := DetialOrder{Detial_order_list:detial_order_list}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "ordersDetial", detial_order)
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