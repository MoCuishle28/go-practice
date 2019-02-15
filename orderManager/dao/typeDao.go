package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
)


func QueryType() *[]entity.Type {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from type")
	checkErr(err)

	var tid string
	var type_name string
	ret := make([]entity.Type, 5)

	for rows.Next() {
		err := rows.Scan(&tid, &type_name)
		checkErr(err)

		t := entity.Type{Tid:tid, Type_name:type_name}

		ret = append(ret, t)
	}
	return &ret
}