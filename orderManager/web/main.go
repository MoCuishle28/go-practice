package main

import (
	"log"
	"os"
	"io"
	"time"

	// 服务器端编程
	"html/template"
	"net/http"

	// util
	"sync"
	"strconv"

	// 自定义 包
	"Go-practice/orderManager/service"
	"Go-practice/orderManager/dao"
	"Go-practice/orderManager/entity"
)


// root 用户登录标志
type Root struct {
	username string  // 未登录时为 "null" 登录后为用户名
	lock sync.Mutex
}

// 以下为模板渲染内容
// 用于渲染登录页模板的信息
type LoginMsg struct {
	Msg string
}

// header
type Header struct {
	Username string
	Index string
}

type Index struct {
	Orders_list *[]entity.Orders
}

type DetialOrder struct {
	Detial_order_list *[]entity.Detial_order
	Status string
}

type DishManage struct {
	Dishes_list *[]entity.Dishes
}

type ActivityManage struct {
	Dish_activity_list *[]entity.Dish_activity
	Order_activity_list *[]entity.Order_activity
	Dishes_list *[]entity.Dishes
}

type DishFrom struct {
	Dish *entity.Dishes
	Types *[]entity.Type
	Add string
}

type SaleStatus struct {
	Days *[]int64
	DaySales *[]float64
	MonthSales *[]float64

	Year_list *[]int64	// 用于下拉框显示

	CurrYear *[]int64
	YearSales *[]float64

	ActiveYear int64
	ActiveDate string

	DishSaleNum_list *[]entity.DishSaleNum
}


var root *Root 			// root 用户登录标志


func init() {
	root = &Root{username:"null"}
}


// 主函数
func main() {
	// 不启动静态文件服务就无法加载 css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务

	http.HandleFunc("/login", login)
	http.HandleFunc("/index", index)
	http.HandleFunc("/ordersdetial", ordersDetial)
	http.HandleFunc("/dishmanage", dishManage)
	http.HandleFunc("/ordermanage", orderManage)
	http.HandleFunc("/activity", activity)
	http.HandleFunc("/finishorder", finishOrder)
	http.HandleFunc("/cancelorder", cancelOrder)
	http.HandleFunc("/minusdishesorders", minusDish_in_Order)
	http.HandleFunc("/dishform", dishForm)
	http.HandleFunc("/addorderactivity", addOrderActivity)
	http.HandleFunc("/adddishactivity", addDishActivity)
	http.HandleFunc("/stoporderactivity", stopOrderActivity)
	http.HandleFunc("/startorderactivity", startOrderActivity)
	http.HandleFunc("/stopdishactivity", stopDsihActivity)
	http.HandleFunc("/startdishactivity", startDishActivity)
	http.HandleFunc("/salestatus", saleStatus)

	err := http.ListenAndServe(":9090", nil)	// 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


// 当前订单页
func index(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	orders_list := dao.QueryCurrentOrders()

	t, _ := template.ParseFiles("templates/index.html", "templates/header.html")
	header := Header{Username:root.username, Index:"0"}
	index := Index{Orders_list:orders_list}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "index", index)
}


// 销售情况
func saleStatus(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}
	var date string
	var year string

	if r.Method == "POST" {
		r.ParseForm()
		date = r.Form.Get("date")
		year = r.Form.Get("year")
	}
	if date == "" {
		date = time.Now().Format("2006-01")
	}
	if year == "" {
		year = "2019"
	}
	
	// service 内部可以优化一下 用 redis 缓存
	days, daySales := service.DaySales(date)
	year_list := service.GetYears()
	monthsale := service.MonthSales(year)
	currYear, yearSales := service.YearSales()
	year_num, _ := 	strconv.ParseInt(year, 10, 64)
	dishSaleNum_list := service.GetDishSaleNum()

	header := Header{Username:root.username, Index:"4"}
	sale_status := SaleStatus{
		MonthSales:monthsale,
		Year_list:year_list,
		CurrYear:currYear,		// 到目前年份的 slice
		YearSales:yearSales,	// 到目前年份的年销售额
		Days:days,
		DaySales:daySales,
		ActiveYear:year_num,
		ActiveDate:date,
		DishSaleNum_list:dishSaleNum_list}

	t, _ := template.ParseFiles("templates/saleStatus.html", "templates/header.html")
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "saleStatus", sale_status)
}


func stopOrderActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "GET" {
		r.ParseForm()
		id := r.Form.Get("id")
		affect := dao.UpdateOrderActivityWork("0", id)
		log.Println("affect:", affect)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


func startOrderActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "GET" {
		r.ParseForm()
		id := r.Form.Get("id")
		affect := dao.UpdateOrderActivityWork("1", id)
		log.Println("affect:", affect)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


func stopDsihActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "GET" {
		r.ParseForm()
		id := r.Form.Get("id")
		affect := dao.UpdateDishActivityWork("0", id)
		log.Println("affect:", affect)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


func startDishActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "GET" {
		r.ParseForm()
		id := r.Form.Get("id")
		affect := dao.UpdateDishActivityWork("1", id)
		log.Println("affect:", affect)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


func addDishActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}
	
	if r.Method == "POST" {
		status := service.AddDishActivity(r)
		log.Println("status:", status)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


func addOrderActivity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "POST" {
		status := service.AddOrderActivity(r)
		log.Println("status:", status)
		http.Redirect(w, r, "/activity", http.StatusFound)
	}
}


// 添加 or 修改 菜品信息
func dishForm(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	if r.Method == "GET" {
		r.ParseForm()
		did := r.Form.Get("did")
		dish := dao.QueryDishByDid(did)
		types := dao.QueryType()

		header := Header{Username:root.username, Index:"2"}
		dishForm := DishFrom{Dish:&dish, Types:types}

		t, _ := template.ParseFiles("templates/dishForm.html", "templates/header.html")
		t.ExecuteTemplate(w, "header", header)
		t.ExecuteTemplate(w, "dishForm", dishForm)
	} else {
		r.ParseMultipartForm(32 << 20)	// 把上传的文件存储在内存和临时文件中

		did := r.FormValue("did")
		name := r.FormValue("name")
		price := r.FormValue("price")
		type_id := r.FormValue("type_id")
		status := r.FormValue("status")
		dish := entity.Dishes{Did:did, Name:name, Price:price, Type_id:type_id, Status:status}

		// 获取文件句柄，然后对文件进行存储等处理
		file, handler, err := r.FormFile("file")
		if err == nil {
			defer file.Close()
			date := time.Now().Format("2006-01-02-15-04-05")
			path := "I:/WeChatMiniSrc/order/pages/images/dish/" + date + handler.Filename
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println("err", err)
				return
			}
			defer f.Close()
			io.Copy(f, file)	// 保存文件	
			path = "../images/dish/" + date + handler.Filename
			dish.Img.String = path
			dish.Img.Valid = true
		}

		var affect int64 = -1
		if did == "" {
			affect = dao.InsertDish(dish)
		} else {
			affect = dao.UpdateDish(dish)
		}
		log.Println("affect:", affect, "post dish:", dish)
		http.Redirect(w, r, "/dishmanage", http.StatusFound)
	}
}


// 删除订单内的菜品
func minusDish_in_Order(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	r.ParseForm()
	oid := r.Form.Get("oid")
	did := r.Form.Get("did")

	ret_status := service.MinusDish_in_Order(oid, did)
	log.Println("status", ret_status)

	http.Redirect(w, r, "/ordersdetial?oid="+oid, http.StatusFound)
}


// 完成订单
func finishOrder(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	r.ParseForm()
	oid := r.Form.Get("oid")
	order := entity.Orders{Oid:oid, Finished:"1"}
	affect := dao.UpdateOrder(&order)
	if affect == -1 {
		log.Println("修改失败")
	} else {
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}


// 撤销订单
func cancelOrder(w http.ResponseWriter, r *http.Request){
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	r.ParseForm()
	oid := r.Form.Get("oid")
	order := entity.Orders{Oid:oid, Finished:"2"}
	affect := dao.UpdateOrder(&order)
	if affect == -1 {
		log.Println("修改失败")
	} else {
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}


// 优惠活动管理页
func activity(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	dish_activity_list := dao.QueryDishActivity()
	order_activity_list := dao.QueryOrderActivity()
	dishes_list := dao.QueryDsihes()
	header := Header{Username:root.username, Index:"1"}
	activity := ActivityManage{Dish_activity_list:dish_activity_list, Order_activity_list:order_activity_list, Dishes_list:dishes_list}

	t, _ := template.ParseFiles("templates/activityManage.html", "templates/header.html")
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "activity", activity)
}


// 菜品管理页
func dishManage(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	dishes_list := dao.QueryDsihes()
	t, _ := template.ParseFiles("templates/dishesManage.html", "templates/header.html")
	header := Header{Username:root.username, Index:"2"}
	dish_manage := DishManage{Dishes_list:dishes_list}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "dishesManage", dish_manage)
}


// 订单管理
func orderManage(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	orders_list := dao.QueryOrders()

	t, _ := template.ParseFiles("templates/index.html", "templates/header.html")
	header := Header{Username:root.username, Index:"3"}
	index := Index{Orders_list:orders_list}
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "index", index)
}


// 订单详情页 GET 请求带参数 ?oid=...
func ordersDetial(w http.ResponseWriter, r *http.Request) {
	root.lock.Lock()
	defer root.lock.Unlock()
	if root.username == "null" {
		login(w, r)
		return
	}

	r.ParseForm()
	oid := r.Form.Get("oid")
	detial_order_list := dao.QueryOrdersDetial(oid)
	order := dao.QueryOrderByOid(oid)

	t, _ := template.ParseFiles("templates/ordersDetial.html", "templates/header.html")
	header := Header{Username:root.username, Index:"0"}
	detial_order := DetialOrder{ Detial_order_list:detial_order_list, Status:order.Finished }
	t.ExecuteTemplate(w, "header", header)
	t.ExecuteTemplate(w, "ordersDetial", detial_order)
}


// 登录页
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {	// POST
		username := r.Form.Get("username")
		pwd := r.Form.Get("pwd")

		msg := LoginMsg{}
		ret := service.ValidLogin(username, pwd)
		switch(ret){
			case -1:
				msg.Msg = "其他错误!"
			case 1:
				root.lock.Lock()
				root.username = username
    			root.lock.Unlock()
    			index(w, r)
    			return
			case 2:
				msg.Msg = "用户名错误!"
			case 3:
				msg.Msg = "密码错误!"
			default:
				msg.Msg = "异常!"
		}
    	t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, msg)
	}
}