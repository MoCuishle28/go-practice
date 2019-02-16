package service


import (
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
	"net/http"
	"log"
	"strconv"
)


func AddOrderActivity(r *http.Request) int {
	r.ParseForm()
	ret_map := map[string]int {"success":1, "error":-1}

	create_time := r.Form.Get("start_time")
	end_time := r.Form.Get("end_time")
	if create_time == "" || end_time == "" {
		log.Println("未填写时间信息")
		return ret_map["error"]
	}

	discount_full := r.Form.Get("discount_full")
	discount := r.Form.Get("discount")
	cost_full_0 := r.Form.Get("cost_full_0")
	minus := r.Form.Get("minus")
	cost_full_1 := r.Form.Get("cost_full_1")
	give := r.Form.Get("give")
	
	activity := entity.Order_activity{Created_time:create_time, End_time:end_time, Work:"1"}

	if discount_full != "" && discount != "" {
		discount_num, _ := strconv.ParseFloat(discount, 64)
		if discount_num >= 1 {
			log.Println("discount 填写错误:", discount)
			return ret_map["error"]
		}
		activity.Discount.String = discount_full+","+discount
		activity.Discount.Valid = true

	} else if cost_full_0 != "" && minus != "" {
		cost_full_0_num, _ := strconv.ParseFloat(cost_full_0, 64)
		minus_num, _ := strconv.ParseFloat(minus, 64)
		if minus_num >= cost_full_0_num {
			log.Println("满减填写错误:", cost_full_0, minus)
			return ret_map["error"]
		}
		activity.Full_minus.String = cost_full_0+","+minus
		activity.Full_minus.Valid = true

	} else if cost_full_1 != "" && give != "" {
		// TODO
	}
	
	affect :=  dao.InsertOrderActivity(&activity)
	if affect == -1 {
		log.Println("SQL 插入信息失败")
		return ret_map["error"]
	}
	return ret_map["success"]
}