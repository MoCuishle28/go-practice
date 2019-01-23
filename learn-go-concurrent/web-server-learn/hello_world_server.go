package main

import (
	"net/http"
	"fmt"
)

func main() {
	// 127.0.0.1:8888/?name=MoCuishle
	http.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>hello world %s</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8888", nil)
}