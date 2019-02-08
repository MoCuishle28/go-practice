package main

import (
	// "fmt"
	"html/template"
	"os"
)

type Person struct{
	UserName string		// 首字母必须大写
}

func main() {
	t := template.New("filename example")
	t, _ = t.Parse("Hello {{.UserName}}")	// {{.}}表示当前的对象
	p := Person{ UserName:"MoCuishle" }
	t.Execute(os.Stdout, p)
}