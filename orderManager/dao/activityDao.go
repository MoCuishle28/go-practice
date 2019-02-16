package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
	"log"
)


func InsertDishActivity(activity *entity.Dish_activity) int64 {
	// insert into dish_activity(discount,work,created_time,end_time) values('','','','')
	if activity.Created_time == "" || activity.End_time=="" || activity.Work==""{
		log.Println("优惠信息未填写完整")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	var sql_str string
	if activity.Discount.Valid {
		sql_str = "insert into dish_activity(did,created_time,end_time,work,discount) values(?,?,?,?,?)"
	} else if activity.Minus_price.Valid {
		sql_str = "insert into dish_activity(did,created_time,end_time,work,minus_price) values(?,?,?,?,?)"
	}

	stmt, err := db.Prepare(sql_str)
	checkErr(err)

	var res sql.Result
	if activity.Discount.Valid {
		res, err = stmt.Exec(activity.Did,activity.Created_time,activity.End_time,activity.Work,activity.Discount.String)
	} else if activity.Minus_price.Valid {
		res, err = stmt.Exec(activity.Did,activity.Created_time,activity.End_time,activity.Work,activity.Minus_price.String)
	}
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


func InsertOrderActivity(activity *entity.Order_activity) int64 {
	// insert into order_activity(discount,work,created_time,end_time) values('','','','')
	if activity.Created_time == "" || activity.End_time=="" || activity.Work==""{
		log.Println("优惠信息未填写完整")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	var sql_str string
	if activity.Discount.Valid {
		sql_str = "insert into order_activity(created_time,end_time,work,discount) values(?,?,?,?)"
	} else if activity.Full_minus.Valid {
		sql_str = "insert into order_activity(created_time,end_time,work,full_minus) values(?,?,?,?)"
	} else if activity.Full_give.Valid {
		sql_str = "insert into order_activity(created_time,end_time,work,full_give) values(?,?,?,?)"
	}

	stmt, err := db.Prepare(sql_str)
	checkErr(err)

	var res sql.Result
	if activity.Discount.Valid {
		res, err = stmt.Exec(activity.Created_time,activity.End_time,activity.Work,activity.Discount.String)
	} else if activity.Full_minus.Valid {
		res, err = stmt.Exec(activity.Created_time,activity.End_time,activity.Work,activity.Full_minus.String)
	} else if activity.Full_give.Valid {
		res, err = stmt.Exec(activity.Created_time,activity.End_time,activity.Work,activity.Full_give.String)
	}
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


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


// 查询 Did 菜品活动
func QueryDishWrokActivityByDid(did string) *[]entity.Dish_activity {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from dish_activity where work='1' and did=?", did)
	checkErr(err)

	var id string
	var work string
	var minus_price sql.NullString
	var discount sql.NullString
	var created_time string
	var end_time string
	ret := make([]entity.Dish_activity, 10)

	for rows.Next() {
		err := rows.Scan(&id, &did, &work, &minus_price, &discount, &created_time, &end_time)
		checkErr(err)

		activity := entity.Dish_activity{Id:id, Work:work, Minus_price:minus_price, Discount:discount, Created_time:created_time, End_time:end_time}
		ret = append(ret, activity)
	}
	return &ret
}


// 查询还能用的菜品活动
func QueryDishWorkActivity() *[]entity.Dish_activity {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from dish_activity where work='1'")
	checkErr(err)

	var id string
	var did string
	var work string
	var minus_price sql.NullString
	var discount sql.NullString
	var created_time string
	var end_time string
	ret := make([]entity.Dish_activity, 10)

	for rows.Next() {
		err := rows.Scan(&id, &did, &work, &minus_price, &discount, &created_time, &end_time)
		checkErr(err)

		activity := entity.Dish_activity{Id:id, Work:work, Minus_price:minus_price, Discount:discount, Created_time:created_time, End_time:end_time}

		ret = append(ret, activity)
	}
	return &ret
}


// 查询还能用的订单活动
func QueryOrderWrokActivity() *[]entity.Order_activity {
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select * from order_activity where work = '1'")
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