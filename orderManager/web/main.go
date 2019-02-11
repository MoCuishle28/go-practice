package main

import (
	"fmt"
	"log"

	// 服务器端编程
	"html/template"
	"net/http"

	// util
	"sync"
	"strings"

	// 用于文件操作
	"os"
	"encoding/json"
	"io/ioutil"
)

type RootLogin struct {
	Username string
	Password string
}

type Root struct {
	username string
	lock sync.Mutex
}


var root *Root
var not_filter_url map[string]bool


func init() {
	root = &Root{username:"null"}
	not_filter_url = make(map[string]bool)
	not_filter_url["favicon.ico"] = true
	not_filter_url["login"] = true
}


// 主函数
func main() {
	http.HandleFunc("/", filter)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)	// 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func filter(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		// 直接返回登录页面
		login(w, r)
		return
	}
	root.lock.Lock()
	defer root.lock.Unlock()

	url := strings.Split(r.URL.String(), "/")
	// 不在不拦截名单的url就要判断是否已经登录
	if _,ok := not_filter_url[url[1]]; !ok {
		if root.username == "null" {
			fmt.Println("未登录 Root...")
			return
		}
	}
	fmt.Println("当前访问用户：", root.username)
}


func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {	// POST 把业务逻辑分离好点吧 TODO
		username := r.Form.Get("username")
		pwd := r.Form.Get("pwd")

		var root_login RootLogin
		file, err := os.Open("root.json")
		if err != nil {
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return
		}

		json.Unmarshal(data, &root_login)
		root_name := root_login.Username
		root_pwd := root_login.Password

    	if root_name == username {
    		if root_pwd == pwd {
    			// TODO 成功登录
    			fmt.Println("success")
    			root.lock.Lock()
    			root.username = root_name
    			root.lock.Unlock()
    		} else {
    			// 密码错误
    			fmt.Println("pwd Error")
    		}
    	} else {
			// 用户名错误
			fmt.Println("name Error")
    	}
	}
}