package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
	"log"
)


func DeleteDishes_OrdersByOidAndDid(dishes_orders *entity.Dishes_orders) int64 {
	// delete from dishes_orders where oid=? and did=?
	if dishes_orders.Oid == "" || dishes_orders.Did == "" {
		log.Println("缺少OID或DID")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	sql := "delete from dishes_orders where oid=? and did=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(dishes_orders.Oid, dishes_orders.Did)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


func UpdateDishes_OrdersByOidAndDid(dishes_orders *entity.Dishes_orders) int64 {
	// update dishes_orders set num = ? where oid=? and did=?
	if dishes_orders.Oid == "" || dishes_orders.Did == "" || dishes_orders.Num == 0 {
		log.Println("缺少OID或DID")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	sql := "update dishes_orders set num = ? where oid = ? and did = ?"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(dishes_orders.Num, dishes_orders.Oid, dishes_orders.Did)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


func QueryDishes_OrdersByOid(oid string) *[]entity.Dishes_orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select do.did,do.num,d.price from dishes_orders as do, dishes as d where do.oid=? and do.did=d.did;", oid)
	checkErr(err)

	var did string
	var num int64
	var price string
	ret := make([]entity.Dishes_orders, 5)

	for rows.Next() {
		err := rows.Scan(&did, &num, &price)
		checkErr(err)

		dishes_orders := entity.Dishes_orders{Oid:oid, Did:did, Num:num, Price:price}
		ret = append(ret, dishes_orders)
	}
	return &ret
}


func QueryDishes_OrdersByOidAndDid(oid, did string) entity.Dishes_orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	row := db.QueryRow("SELECT num FROM dishes_orders where did=? and oid=?", did, oid)

	var num int64
	err = row.Scan(&num)
	checkErr(err)

	dishes_orders := entity.Dishes_orders{Oid:oid, Did:did, Num:num}
	return dishes_orders
}