package router


import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"Go-practice/SecKill/SecProxy/controller"
)


func init() {
	logs.Debug("start router init")
	// 路由绑定 1.路由 2.控制 3.具体处理函数（*代表既可以 GET 也可以 POST）
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}