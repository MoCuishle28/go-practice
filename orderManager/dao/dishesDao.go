package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
)


func QueryDishByDid(did string) entity.Dishes {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	row := db.QueryRow("SELECT * FROM dishes where did=?", did)	// 期待返回一条结果

	var name string
	var price string
	var img sql.NullString
	var type_id string
	var status string

	err = row.Scan(&did, &name, &price, &img, &type_id, &status)
	checkErr(err)

	dish := entity.Dishes{Did:did, Name:name, Price:price, Img:img, Type_id:type_id, Status:status}
	return dish
}


func QueryDsihes() *[]entity.Dishes {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT d.did,d.name,d.price,d.img,d.type_id,t.type,d.status FROM dishes as d,type as t where d.type_id=t.tid")
	checkErr(err)

	var did string
	var name string
	var price string
	var img sql.NullString
	var type_id string
	var type_name string
	var status string
	ret_dishes := make([]entity.Dishes, 10)

	for rows.Next() {
		err := rows.Scan(&did, &name, &price, &img, &type_id, &type_name, &status)
		checkErr(err)

		dish := entity.Dishes{Did:did, Name:name, Price:price, Img:img, Type_id:type_id, Type_name:type_name, Status:status}

		ret_dishes = append(ret_dishes, dish)
	}
	return &ret_dishes
}