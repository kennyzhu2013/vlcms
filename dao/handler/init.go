/*
@Time : 2018/8/16 15:26 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package handler

import (
	example "github.com/kennyzhu2013/go-os/src/dbservice/proto/example"
	"github.com/micro/go-micro/server"
)

//all handlers init here...
func Init(s server.Server) {
	//register handler here, each request run in goroutine...
	//example.RegisterPreferencesHandler(service.Server(), new(handler.Example))
	example.RegisterPreferencesHandler(s, new(Example))

	//To register others...
}