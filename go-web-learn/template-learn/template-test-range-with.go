package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`
			hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}

			{{$a := "minux.ma"}}	<!-- 要先定义变量-->
			{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}

					<!-- if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com" -->
					{{ if eq .Fname $a}}	<!-- 用上了 eq 函数 -->
						yes!!! .Fname eq $a
					{{end}}

				{{end}}
			{{end}}
			`)

	p := Person{
		UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}

	t.Execute(os.Stdout, p)
}