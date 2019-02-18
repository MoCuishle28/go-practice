package service

import (
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
	"Go-practice/orderManager/util"
	"strings"
	"strconv"
)


var month_day_map = map[int]int{1:31,2:28,3:31,4:30,5:31,6:30,7:31,8:31,9:30,10:31,11:30,12:31}


func changeMonth() {
	month_day_map[2] = 28
}


func GetYears() *[]int64 {
	tmp_arr := strings.Split(dao.QueryMinDate(), " ")
	date_arr := strings.Split(tmp_arr[0], "-")
	min_date, _ := strconv.ParseInt(date_arr[0], 10, 0)
	var year []int64
	for i := min_date; i <= min_date+50; i++ {
		year = append(year, i)
	}
	return &year
}


func DaySales(date string) (*[]int64, *[]float64){
	date_arr := strings.Split(date, "-")
	year, month := date_arr[0], date_arr[1]
	month_num, _ := strconv.ParseInt(month, 10, 64)
	var days []int64
	var daySales []float64
	var end_date string

	if month_num < 10 {
		end_date = year + "-0" + strconv.FormatInt(month_num+1, 10)
	} else {
		end_date = year + "-" + strconv.FormatInt(month_num+1, 10)
	}
	orders_list := dao.QueryOrdersByDate(date, end_date)

	if util.IsLeapYear(year) && month_num == 2 {
		month_day_map[2] = 29
		defer changeMonth()
	}
	for i := 0; i < month_day_map[int(month_num)]; i++ {
	 	days = append(days, int64(i+1))
	 	daySales = append(daySales, 0.0)
	}

	for _, v := range *orders_list {
		if v.Oid == "" {
			continue
		}
		tmp_arr := strings.Split(v.Created_time, " ")
		arr := strings.Split(tmp_arr[0], "-")
		index, _ := strconv.ParseInt(arr[2], 10, 0)
		final_cost, _ := strconv.ParseFloat(v.Final_cost, 64)
		daySales[index] += final_cost
	}
	return &days, &daySales
}


func MonthSales(year string) *[]float64 {
	monthSlaes := make([]float64, 12)
	year_num, _ := strconv.ParseInt(year, 10, 0)
	endYear := strconv.FormatInt(year_num+1, 10)
	orders_list := dao.QueryOrdersByDate(year, endYear)
	for _, v := range *orders_list {
		if v.Oid == "" {
			continue
		}
		tmp_arr := strings.Split(v.Created_time, " ")
		arr := strings.Split(tmp_arr[0], "-")
		month, _ := strconv.ParseInt(arr[1], 10, 0)
		final_cost, _ := strconv.ParseFloat(v.Final_cost, 64)
		monthSlaes[month-1] += final_cost
	}
	return &monthSlaes
}


func YearSales() (*[]int64, *[]float64) {
	min_date := dao.QueryMinDate()
	max_date := dao.QueryMaxDate()
	tmp_arr0, tmp_arr1 := strings.Split(min_date, " "), strings.Split(max_date, " ")
	date_arr0, date_arr1 := strings.Split(tmp_arr0[0], "-"), strings.Split(tmp_arr1[0], "-")
	endYear_num, _ := strconv.ParseInt(date_arr1[0], 10, 0)
	startYear, endYear := date_arr0[0], strconv.FormatInt(endYear_num+1, 10)

	orders_list := dao.QueryOrdersByDate(startYear, endYear)

	startYear_num, _ := strconv.ParseInt(startYear, 10, 0)
	currYear := make([]int64, endYear_num - startYear_num + 1)
	yearSales := make([]float64, endYear_num - startYear_num + 1)

	year := startYear_num
	year_index_map := map[int]int{}
	for i,_ := range currYear {
		currYear[i] = year
		yearSales[i] = 0.0
		year_index_map[int(year)] = i
		year++
	}

	for _, v := range *orders_list {
		if v.Oid == "" {
			continue
		}
		tmp_arr := strings.Split(v.Created_time, " ")
		arr := strings.Split(tmp_arr[0], "-")
		year, _ := strconv.ParseInt(arr[0], 10, 0)
		final_cost, _ := strconv.ParseFloat(v.Final_cost, 64)
		yearSales[year_index_map[int(year)]] += final_cost
	}

	return &currYear, &yearSales
}


func GetDishSaleNum() *[]entity.DishSaleNum {
	dishes_orders_list := dao.QueryDishSalesNum()
	dishSaleNum_list := make([]entity.DishSaleNum, 5)
	for _, v := range *dishes_orders_list {
		if v.Name == "" {
			continue
		}
		dishSaleNum_list = append(dishSaleNum_list, entity.DishSaleNum{Name:v.Name, Value:v.Num})
	}
	return &dishSaleNum_list
}