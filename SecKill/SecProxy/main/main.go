package main


import(
	"github.com/astaxie/beego"
	_ "Go-practice/SecKill/SecProxy/router"
)



func main() {
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return 
	}
	beego.Run()
}