package service

import (
	"gopkg.in/macaron.v1"
	"goweb/route/adminRoute"
)

func init() {
	initEnv()        //初始化环境变量等
}

func Run() {
	m := macaron.Classic()

	initStaticDir(m)
	initMiddleWare(m)

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "saplmm"
		ctx.HTML(200, "index") // 200 为响应码
	})

	m.Get("/golang", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "saplmm-golang"
		ctx.HTML(200, "golang") // 200 为响应码
	})

	route.AdminRout(m)        //后台的路由
	m.Run(8080)
}

func initEnv() {
	macaron.Env = macaron.PROD        //设置环境变量
}

func initStaticDir(m *macaron.Macaron) {
	//m.Use(macaron.Static("public"))	//这里可以不用设置,因为这个是默认值
}

func initMiddleWare(m *macaron.Macaron) {
	m.Use(macaron.Renderer())                //模板引擎

}