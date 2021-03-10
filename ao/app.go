/*
@Time : 2018/8/30 16:01
@Author : kenny zhu
@File : app.go
@Software: GoLand
@Others:
*/
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kennyzhu2013/vlcms/ao/Controller"
)

var App *iris.Application

func Start(runner iris.Runner) {
	App = iris.New()

	App.Logger().SetLevel("debug")
	App.Use(recover.New())
	App.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(exit)
	})

	// views register..
	App.RegisterView(iris.HTML("./web/views", ".html"))

	// root path..
	App.Get("/", func(ctx iris.Context) {
		ctx.Exec("GET", "/index") // like it was called by the client.
	})
	mvc.Configure(App.Party("/index"), Controller.Index)

	_ = App.Run(runner,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations)
}

func exit() {
	println("server terminated!")
}
