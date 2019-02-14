package util

import (
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
	"strconv"
	"strings"
)


// 计算出按照菜品优惠的最大优惠
func ChooseDishActivity(order *entity.Orders, dishes_orders_list *[]entity.Dishes_orders) entity.Orders {
	ret := entity.Orders{Oid:order.Oid, Original_cost:order.Original_cost, Oa_id:"-1"}
	var final_cost float64 = 0.0
    var dish_cost float64 = 9999999

    for _, dish := range *dishes_orders_list {
        dish_activities := dao.QueryDishWrokActivityByDid(dish.Did)
        price, _ := strconv.ParseFloat(dish.Price, 64)
        dish_cost = price * float64(dish.Num)
        for _, activity := range *dish_activities {
            var tmp_cost float64 = 0.0
            if activity.Discount.Valid != false {
            	discount, _ := strconv.ParseFloat(activity.Discount.String, 64)
                tmp_cost = (price * discount) + (price * float64(dish.Num - 1))
            } else {
                // TODO 这里似乎有 BUG (2*牛肉串,2*羊肉串,2*百合酱蒸凤爪 百合减一时出错)
            	minus_price, _ := strconv.ParseFloat(activity.Minus_price.String, 64)
                tmp_cost = (price - minus_price) * float64(dish.Num)
            }

            if tmp_cost < dish_cost {
            	dish_cost = tmp_cost
            }
        }
        final_cost += dish_cost;
    }

    ret.Final_cost = strconv.FormatFloat(final_cost, 'f', 2, 64)
    return ret
}


// 计算出按照订单优惠的最大优惠
func ChooseOrderActivity(order *entity.Orders) entity.Orders {
	ret := entity.Orders{Oid:order.Oid, Original_cost:order.Original_cost, Oa_id:order.Oa_id}
	order_activities := dao.QueryOrderWrokActivity()
    var use_oa_id string = "-1"   		// 用到的优惠活动 id
    original_cost, _ := strconv.ParseFloat(order.Original_cost, 64)
    cost_by_order_activity, _ := strconv.ParseFloat(order.Original_cost, 64)
    var tmp_order_cost float64 = 9999999

    for _, o := range *order_activities {
        if o.Discount.Valid != false {
            str_arr := strings.Split(o.Discount.String, ",")
            full, _ := strconv.ParseFloat(str_arr[0], 64)
            discount, _ := strconv.ParseFloat(str_arr[1], 64)
            if original_cost >= full {
                tmp_order_cost = original_cost * discount
            }
        } else if o.Full_minus.Valid != false {
            str_arr := strings.Split(o.Full_minus.String, ",")
            full, _ := strconv.ParseFloat(str_arr[0], 64)
            minus, _ := strconv.ParseFloat(str_arr[1], 64)
            if original_cost >= full {
                tmp_order_cost = original_cost - minus;
            }
        }

        if tmp_order_cost < cost_by_order_activity {
        	use_oa_id = o.Id
        	cost_by_order_activity = tmp_order_cost
        }
        tmp_order_cost = 9999999
    }

    ret.Oa_id = use_oa_id
    ret.Final_cost = strconv.FormatFloat(cost_by_order_activity, 'f', 2, 64)
    return ret
}