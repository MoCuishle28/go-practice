package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
)


func QueryDishSalesNum() *[]entity.Dishes_orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	row, err := db.Query("select d.name,sum(do.num) from dishes_orders as do, dishes as d where do.did=d.did group by do.did")
	checkErr(err)

	var num int64
	var name string
	ret := make([]entity.Dishes_orders, 10)

	for row.Next() {
		err = row.Scan(&name,&num)
		checkErr(err)
		dishes_orders := entity.Dishes_orders{Name:name, Num:num}
		ret = append(ret, dishes_orders)
	}

	return &ret
}