package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func def(a int) {
	fmt.Println("func def()",a)
}

func add_a(a *int) int {
	*a = *a + 1
	fmt.Println("add_a:", *a)
	return *a
}

func c() {
	a := 9
	defer def(add_a(&a))
	fmt.Println("func c():",a)
	fmt.Println("func c() end")
}

func main() {
	b()
	fmt.Println("-------------------")
	c()
}