package main

import "gopkg.in/macaron.v1"

func init() {
	initEnv()        //初始化环境变量等
}

func main() {
	m := macaron.Classic()

	initStaticDir(m)
	initMiddleWare(m)

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "saplmm"
		ctx.HTML(200, "index") // 200 为响应码
	})
	m.Run()
}

func initEnv() {
	//macaron.Env = macaron.PROD	//设置环境变量
}

func initStaticDir(m *macaron.Macaron) {
	//m.Use(macaron.Static("public"))	//这里可以不用设置,因为这个是默认值
}

func initMiddleWare(m *macaron.Macaron) {
	m.Use(macaron.Renderer())                //模板引擎

}