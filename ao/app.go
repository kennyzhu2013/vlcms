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
	"github.com/kataras/iris/mvc"
	"github.com/kennyzhu/vlcms/ao/Controller"
)

var App *iris.Application


func Start(runner iris.Runner)  {
	App = iris.New()

	App.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(exit)
	})

	// root path..
	mvc.Configure( App.Party("/"), Controller.Init )

	App.Run( runner )
}

func exit() {
	println("server terminated!")
}
