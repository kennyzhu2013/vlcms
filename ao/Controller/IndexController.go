/*
@Time : 2018/8/30 16:28 
@Author : kenny zhu
@File : IndexController.go
@Software: GoLand
@Others:
*/
package Controller

import "github.com/kataras/iris/mvc"

type IndexController struct {
	// mvc.BaseController

	visits uint64 // all request share data.
}

// BeforeActivation called once before the server ran, and before the routes and dependencies binded.
func (m *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// self define route
	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	b.Handle("GET", "/index/{id:long}", "MyCustomHandler", nil)
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080, the same as /index
func (m *IndexController) Get() string { return m.GetIndex() }

// Get serves
// Method:   GET
// Resource: http://localhost:8080/index
func (m *IndexController) GetIndex() string { return "Hey" }

// GET: http://localhost:8080/root/something/{id:long}
func (m *IndexController) OpenGame(id int64) string { return "MyCustomHandler says Hey" }