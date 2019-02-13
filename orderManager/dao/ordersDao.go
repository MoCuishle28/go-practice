package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
	"Go-practice/orderManager/util"
	"log"
	"strings"
)


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func QueryOrderByOid(oid string) entity.Orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows := db.QueryRow("SELECT * FROM orders where oid=?", oid)

	var original_cost string
	var final_cost string
	var finished string
	var created_time string
	var oa_id string
	var uid string

	err = rows.Scan(&oid, &original_cost, &final_cost, &finished, &created_time, &oa_id, &uid)
	checkErr(err)

	order := entity.Orders{Oid:oid, Original_cost:original_cost, Final_cost:final_cost,
		Finished:finished, Created_time:created_time, Oa_id:oa_id, Uid:uid}

	return order
}


func UpdateOrder(order *entity.Orders) int64 {
	// update orders set work='1' where oid=?
	if order.Oid == "" {
		log.Println("缺少OID")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	sql := "update orders set "
	set_arr := make([]string, 6)

	if order.Original_cost != "" {
		set_arr = append(set_arr, "original_cost="+order.Original_cost)
	}
	if order.Final_cost != "" {
		set_arr = append(set_arr, "final_cost="+order.Final_cost)
	}
	if order.Finished != "" {
		set_arr = append(set_arr, "finished="+order.Finished)
	}
	if order.Created_time != "" {
		set_arr = append(set_arr, "created_time="+order.Created_time)
	}
	if order.Oa_id != "" {
		set_arr = append(set_arr, "oa_id="+order.Oa_id)
	}
	if order.Uid != "" {
		set_arr = append(set_arr, "uid="+order.Uid)
	}

	set_arr = util.RemoveZero(set_arr)
	sql += strings.Join(set_arr, ",")
	sql += " where oid = ?"

	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(order.Oid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


func QueryOrdersDetial(oid string) *[]entity.Detial_order {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	sql := "select d.did,d.name,d.price,do.num,t.type,do.oid from dishes_orders as do, orders as o, dishes as d, type as t where o.oid=do.oid and do.did=d.did and o.oid=? and d.type_id=t.tid"
	rows, err := db.Query(sql, oid)
	checkErr(err)

	var did string
	var name string
	var price string
	var num  string
	var type_name string
	ret := make([]entity.Detial_order, 10)

	for rows.Next() {
		err := rows.Scan(&did, &name, &price, &num, &type_name, &oid)
		checkErr(err)

		detial := entity.Detial_order{Did:did, Name:name, Price:price, Num:num, Type_name:type_name, Oid:oid}
		ret = append(ret, detial)
	}
	return &ret
}


func QueryOrders() *[]entity.Orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM orders")		// 直接执行Sql返回Rows结果
	checkErr(err)

	var oid string
	var original_cost string
	var final_cost string
	var finished string
	var created_time string
	var oa_id string
	var uid string
	ret_orders := make([]entity.Orders, 10)

	for rows.Next() {
		err := rows.Scan(&oid, &original_cost, &final_cost, &finished, &created_time, &oa_id, &uid)
		checkErr(err)

		order := entity.Orders{Oid:oid, Original_cost:original_cost, Final_cost:final_cost,
			Finished:finished, Created_time:created_time, Oa_id:oa_id, Uid:uid}

		ret_orders = append(ret_orders, order)
	}
	return &ret_orders
}


func QueryCurrentOrders() *[]entity.Orders {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from orders where created_time >= (select curdate()) order by created_time desc")
	checkErr(err)

	var oid string
	var original_cost string
	var final_cost string
	var finished string
	var created_time string
	var oa_id string
	var uid string
	ret_orders := make([]entity.Orders, 10)

	for rows.Next() {
		err := rows.Scan(&oid, &original_cost, &final_cost, &finished, &created_time, &oa_id, &uid)
		checkErr(err)

		order := entity.Orders{Oid:oid, Original_cost:original_cost, Final_cost:final_cost,
			Finished:finished, Created_time:created_time, Oa_id:oa_id, Uid:uid}

		ret_orders = append(ret_orders, order)
	}
	return &ret_orders
}