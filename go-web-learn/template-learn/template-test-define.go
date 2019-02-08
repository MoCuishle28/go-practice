package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")	// 把所有的嵌套模板全部解析到模板里面

	// 当执行 s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西
	// 通过 ExecuteTemplate(...) 来执行相应的子模板内容
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()

	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()

	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()

	s1.Execute(os.Stdout, nil)
}