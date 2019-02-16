package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"Go-practice/orderManager/entity"
	"log"
	"strings"
)


func InsertDish(dish entity.Dishes) int64 {
	// insert into dishes(did,name,price,type_id,status) values(,'','','','')
	if dish.Name == "" || dish.Price=="" || dish.Type_id=="" || dish.Status=="" {
		log.Println("菜品信息未填写完整")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	var sql_str string
	if dish.Img.Valid == false {
		sql_str = "insert into dishes(name,price,type_id,status) values(?,?,?,?)"
	} else {
		sql_str = "insert into dishes(name,price,type_id,status,img) values(?,?,?,?,?)"
	}

	stmt, err := db.Prepare(sql_str)
	checkErr(err)

	var res sql.Result
	if dish.Img.Valid == false {
		res, err = stmt.Exec(dish.Name,dish.Price,dish.Type_id,dish.Status)
	} else {
		res, err = stmt.Exec(dish.Name,dish.Price,dish.Type_id,dish.Status, dish.Img.String)
	}
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


func UpdateDish(dish entity.Dishes) int64 {
	// update dishes set price='1' where did=?
	if dish.Did == "" {
		log.Println("缺少DID")
		return -1
	}
	db, err := sql.Open("mysql", "test:123456@/wechat_applets?charset=utf8")
	checkErr(err)
	defer db.Close()

	sql := "update dishes set "
	set_arr := make([]string, 5)

	if dish.Name != "" {
		set_arr = append(set_arr, "name='"+dish.Name+"'")
	}
	if dish.Price != "" {
		set_arr = append(set_arr, "price='"+dish.Price+"'")
	}
	if dish.Img.Valid != false {
		set_arr = append(set_arr, "img='"+dish.Img.String+"'")
	}
	if dish.Type_id != "" {
		set_arr = append(set_arr, "type_id="+dish.Type_id)
	}
	if dish.Status != "" {
		set_arr = append(set_arr, "status='"+dish.Status+"'")
	}

	set_arr = removeZero(set_arr)
	sql += strings.Join(set_arr, ",")
	sql += " where did = ?"

	log.Println("SQL: ", sql)

	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(dish.Did)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect
}


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

	_ = row.Scan(&did, &name, &price, &img, &type_id, &status)

	dish := entity.Dishes{Did:did, Name:name, Price:price, Img:img, Type_id:type_id, Status:status}
	log.Println(dish)
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