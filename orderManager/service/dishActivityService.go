package service


import (
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
	"net/http"
	"log"
	"strconv"
)


func AddDishActivity(r *http.Request) int {
	r.ParseForm()
	ret_map := map[string]int {"success":1, "error":-1}

	did := r.Form.Get("did")
	create_time := r.Form.Get("start_time")
	end_time := r.Form.Get("end_time")
	if create_time == "" || end_time == "" || did == "" {
		log.Println("未填写时间信息")
		return ret_map["error"]
	}

	discount := r.Form.Get("discount")
	minus := r.Form.Get("minus")
	
	activity := entity.Dish_activity{Did:did, Created_time:create_time, End_time:end_time, Work:"1"}

	if discount != "" {
		discount_num, _ := strconv.ParseFloat(discount, 64)
		if discount_num >= 1 {
			log.Println("discount 填写错误:", discount)
			return ret_map["error"]
		}
		activity.Discount.String = discount
		activity.Discount.Valid = true

	} else if minus != "" {
		dish := dao.QueryDishByDid(did)
		minus_num, _ := strconv.ParseFloat(minus, 64)
		price, _ := strconv.ParseFloat(dish.Price, 64)
		if minus_num <= 0 || minus_num >= price {
			log.Println("减价填写错误:", "原价:",price, "减价:",minus)
			return ret_map["error"]
		}
		activity.Minus_price.String = minus
		activity.Minus_price.Valid = true
	}

	affect :=  dao.InsertDishActivity(&activity)
	if affect == -1 {
		log.Println("SQL 插入信息失败")
		return ret_map["error"]
	}
	return ret_map["success"]
}