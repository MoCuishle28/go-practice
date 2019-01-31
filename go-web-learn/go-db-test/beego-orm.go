package main

import (
	// "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}


func init() {
	//设置默认数据库
	//mysql用户：test 密码:123456 ， 数据库名称:learn_go_db ， 数据库别名:default
	orm.RegisterDataBase("default", "mysql", "test:123456@/learn_go_db?charset=utf8", 30)

	//注册定义的model
    orm.RegisterModel(new(User))

	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

   	// 创建table
	orm.RunSyncdb("default", false, true)
}


func main() {
    o := orm.NewOrm()	// 开到数据库的链接，然后创建一个beego orm对象

    user := User{Name: "slene"}

    // 插入表
    id, err := o.Insert(&user)
    fmt.Println("插入表")
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // 更新表
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Println("更新表 user : ", user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // 读取 one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Println("读取 one ->u: ", u)
    fmt.Printf("ERR: %v\n", err)

    // 删除表？  还是删除u的数据？
    num, err = o.Delete(&u)
    fmt.Println("删除 u 那条数据")
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}