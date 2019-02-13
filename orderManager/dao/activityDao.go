package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
)


func QueryDishActivity() *[]entity.Dish_activity {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select da.id,da.did,d.name,da.work,da.minus_price,da.discount,da.created_time,da.end_time from dish_activity as da, dishes as d where da.did=d.did")
	checkErr(err)

	var id string
	var did string
	var name string
	var work string
	var minus_price sql.NullString
	var discount sql.NullString
	var created_time string
	var end_time string
	ret := make([]entity.Dish_activity, 10)

	for rows.Next() {
		err := rows.Scan(&id, &did, &name, &work, &minus_price, &discount, &created_time, &end_time)
		checkErr(err)

		activity := entity.Dish_activity{Id:id, Name:name, Work:work, Minus_price:minus_price, Discount:discount, Created_time:created_time, End_time:end_time}

		ret = append(ret, activity)
	}
	return &ret
}


func QueryOrderActivity() *[]entity.Order_activity {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from order_activity")
	checkErr(err)

	var id string
	var discount sql.NullString
	var full_minus sql.NullString
	var full_give sql.NullString
	var work string
	var created_time string
	var end_time string
	ret := make([]entity.Order_activity, 10)

	for rows.Next() {
		err := rows.Scan(&id, &discount, &full_minus, &full_give, &work, &created_time, &end_time)
		checkErr(err)

		activity := entity.Order_activity{Id:id, Discount:discount, Full_minus:full_minus, Full_give:full_give, Work:work, Created_time:created_time, End_time:end_time}

		ret = append(ret, activity)
	}
	return &ret
}