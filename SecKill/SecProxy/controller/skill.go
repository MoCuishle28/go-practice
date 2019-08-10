package controller

import(
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"Go-practice/SecKill/SecProxy/service"
)


type SkillController struct {
	beego.Controller
}


// 秒杀功能
// 处理接口 /seckill?product=20&source=android&authcode=xx&time=xx&nance=xx
func (p *SkillController) SecKill() {
	productId, err := p.GetInt("product_id")

	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"

	defer func() {
		// 写入自动生成 json 格式
		p.Data["json"] = result
		p.ServeJSON() 		// 返回 json 格式数据给前端
	} ()

	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = "invalid product_id"
		return 
	}

	// 获取各个参数
	source := p.GetString("src")
	authcode := p.GetString("authcode")		// 权限
	secTime := p.GetString("time")
	nance := p.GetString("nance")			// 随机数

	userId, err := strconv.Atoi(p.Ctx.GetCookie("userid"))
	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = err.Error()
	}

	secRequest := &service.SecRequest{
		Source: source,
		AuthCode: authcode,
		Nance: nance,
		ProductId: int64(productId),
		SecTime: secTime,
		UserAuthSign: p.Ctx.GetCookie("userAuthSign"),
		UserId: int64(userId),
		AccessTime: time.Now(),
	}


	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		return 
	}

	result["data"] = data
	result["code"] = code

	return
}


func (p *SkillController) SecInfo() {
	// 客户端传过来的参数 分为 Int 和 String
	productId, err := p.GetInt("product_id")

	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"
	// 提取公共代码
	defer func() {
		// 写入自动生成 json 格式
		p.Data["json"] = result
		p.ServeJSON() 		// 返回 json 格式数据给前端
	} ()

	// 没有传入商品 ID 参数 则返回商品列表
	if err != nil {
		data, code, err := service.SecInfoList()
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("service.SecInfoList failed, err:%v", err)
			return
		}
		result["code"] = code
		result["data"] = data
	} else { 	// 否则返回具体商品信息
		data, code, err := service.SecInfo(productId)
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()
			logs.Error("service.SecInfo failed, err:%v", err)
			return
		}
		result["code"] = code
		result["data"] = data
	}
}