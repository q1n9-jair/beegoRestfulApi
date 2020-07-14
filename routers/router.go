package routers

import (
	"beeApi/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//验证登录态过滤器
var FilterUser = func(ctx *context.Context) {
	var uid = ctx.Input.Session("uid")
	if uid == nil {
		res := &controllers.ResultMsg{
			Code: 302,
			Msg:  "not user login",
		}
		ctx.Output.JSON(res, true, true)
	}

}

func init() {
	// 验证用户是否已经登录
	beego.InsertFilter("bee/v1/*", beego.BeforeRouter, FilterUser)
	//需要登录
	notFilterUser := beego.NewNamespace("/bee",
		beego.NSRouter("/v1/getuser", &controllers.UserController{}, "post:GetUser"),
	)
	//免登录模块
	ns := beego.NewNamespace("/bee",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/regin", &controllers.UserController{}, "post:Regin"),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(notFilterUser)
}
