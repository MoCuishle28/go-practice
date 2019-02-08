package main

import (
	"fmt"
	"log"
	"net/http"
	// "time"
)


func main() {
	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func setCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// expiration := time.Now()
	// expiration = expiration.AddDate(1, 0, 0)	// 过期时间

	// fmt.Println("expiration:", expiration)

	name := r.Form.Get("name")
	value := r.Form.Get("value")
	// cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	cookie := http.Cookie{Name: name, Value: value, MaxAge:20}

	fmt.Println(cookie)

	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "成功添加cookie!")
}


func getCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	cookie, _ := r.Cookie(name)
	// cookie.MaxAge = -1	// 删除再写入 response 会产生空指针错误？
	fmt.Fprint(w, cookie)

	// 另一种读取cookie方式
	for _, cookie := range r.Cookies() {
		fmt.Println(cookie)
	}
}