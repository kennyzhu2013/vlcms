/*
@Time : 2018/8/30 16:23 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others:
*/
package Controller

import "github.com/kataras/iris/mvc"

func Init(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle( &IndexController{0} )
}

