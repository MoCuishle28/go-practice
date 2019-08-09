package controller

import(
	"github.com/astaxie/beego"
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
	p.Data["json"] = "sec info"
	p.ServeJSON()
}