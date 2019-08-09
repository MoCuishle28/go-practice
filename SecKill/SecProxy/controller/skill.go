package controller

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"Go-practice/SecKill/SecProxy/service"
)


type SkillController struct {
	beego.Controller
}


func (p *SkillController) SecKill() {
	// 写入自动生成 json 格式
	p.Data["json"] = "sec kill"
	// 返回json数据
	p.ServeJSON()
}


func (p *SkillController) SecInfo() {
	// 客户端传过来的参数 分为 Int 和 String
	productId, err := p.GetInt("product_id")

	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"
	// 提取公共代码
	defer func() {
		p.Data["json"] = result
		p.ServeJSON() 		// 返回 json 格式数据给前端
	} ()

	if err != nil {
		result["code"] = 1001 	// 自己约定 方便前端处理
		result["message"] = "invalid product_id"
		logs.Error("invalid request, get product_id failed, err:%v", err)
		return
	}

	data, code, err := service.SecInfo(productId)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		logs.Error("service.SecInfo failed, err:%v", err)
		return
	}
	result["data"] = data
}