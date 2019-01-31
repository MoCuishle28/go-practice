package main

import (
	"database/sql"
	"fmt"
	"reflect"	// 用来看变量类型的
	_ "github.com/go-sql-driver/mysql"
)

/**
用到的数据表：userinfo
*/

func showRows(rows *sql.Rows) {
	var uid int
	var username string
	var department string
	var created string
	fmt.Println("--------------show begin--------------------")
	for rows.Next() {
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}
	fmt.Println("--------------show end--------------------")
}


func main() {
	// sql.Open() 打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动
	// 第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息
	db, err := sql.Open("mysql", "test:123456@/learn_go_db?charset=utf8")
	checkErr(err)

	//插入数据
	// 返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	// 参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入
	stmt, err := db.Prepare("insert userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")	// 执行stmt准备好的SQL语句
	checkErr(err)

	id, err := res.LastInsertId()	// 最后一条插入的数据的主键 id
	checkErr(err)

	fmt.Println("insert:",id)
	// 看一下插入后的数据表
	rows, err := db.Query("SELECT * FROM userinfo")		// 直接执行Sql返回Rows结果
	checkErr(err)	
	showRows(rows)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)		// 执行stmt准备好的SQL语句
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("update affect:", affect)
	// 看一下更新后的数据表
	rows, err = db.Query("SELECT * FROM userinfo")		// 直接执行Sql返回Rows结果
	checkErr(err)
	showRows(rows)

	//查询数据
	rows, err = db.Query("SELECT * FROM userinfo")		// 直接执行Sql返回Rows结果
	checkErr(err)

	fmt.Println("rows type:", reflect.TypeOf(rows))		// 查看变量类型
	showRows(rows)

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("delete affect:", affect)
	// 删除后的数据表
	rows, err = db.Query("SELECT * FROM userinfo")		// 直接执行Sql返回Rows结果
	checkErr(err)
	showRows(rows)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}