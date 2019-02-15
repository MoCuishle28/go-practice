package service

import (
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
	"Go-practice/orderManager/util"
	"strconv"
	"log"
)


func MinusDish_in_Order(oid, did string) int64 {
	ret_map := map[string]int64 { "success":1, "error":-1 }
	order := dao.QueryOrderByOid(oid)
	dishes_orders_list := dao.QueryDishes_OrdersByOid(oid)
	dishes_orders := entity.Dishes_orders{Oid:oid, Did:did}

	// 菜品数减一
	for i, v := range *dishes_orders_list {
		if v.Did == did && v.Oid == oid {
			(*dishes_orders_list)[i].Num -= 1
			dishes_orders.Num = (*dishes_orders_list)[i].Num
			break
		}
	}

	// 计算减一后的原始总价
	var original_cost float64 = 0.0
	for _, v := range *dishes_orders_list {
		for  i := 0; int8(i) < v.Num; i++ {
			cost, _ := strconv.ParseFloat(v.Price, 64)
			original_cost += cost
		}
	}
	order.Original_cost = strconv.FormatFloat(original_cost, 'f', 2, 64)

	// 更新 菜品_订单
	var affect int64 = 0
	if dishes_orders.Num > 0 {	// 大于零 直接更新
		affect = dao.UpdateDishes_OrdersByOidAndDid(&dishes_orders)
	} else {					// 等于零 删除记录
		affect = dao.DeleteDishes_OrdersByOidAndDid(&dishes_orders)
	}

	log.Println("Original_cost:", order.Original_cost, " dishes_orders:", dishes_orders)
	log.Println("affect:", affect)
	if affect == -1 {
		return ret_map["error"]
	}

	// 计算并更新订单最终价
	order_by_dish_activity := util.ChooseDishActivity(&order, dishes_orders_list)
	order_by_order_activity := util.ChooseOrderActivity(&order)

	if order_by_dish_activity.Final_cost < order_by_order_activity.Final_cost {
		order = order_by_dish_activity
	} else {
		order = order_by_order_activity
	}

	log.Println(order)

	affect = dao.UpdateOrder(&order)	// 订单没菜时也不能删 后面应该还能随时加
	if affect == -1 {
		return ret_map["error"]
	}
	return ret_map["success"]
}