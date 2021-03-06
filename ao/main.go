/*
@Time : 2018/8/30 15:56
@Author : kenny zhu
@File : main.go
@Software: GoLand
@Others:
*/
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kennyzhu2013/vlcms/conf"
	"github.com/micro/go-micro/registry"
	"github.com/pborman/uuid"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	service = &registry.Service{
		Name: "go.micro.srv.aoiris",
		Nodes: []*registry.Node{
			{
				Id:      "go.micro.srv.aoiris-" + uuid.NewUUID().String(),
				Address: "127.0.0.1:8400",
			},
		},
	}
)

// This is example for api gates..
// To delete....
func main() {
	conf.Init("./conf/aoiris.json")
	service.Name = conf.AppConf.ApiName

	// register call back.
	App.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	})

	address := conf.AppConf.IpAddress + ":" + strconv.Itoa(int(conf.AppConf.Port))
	go Start(iris.Addr(address)) // eg:":8400"
	// Register modules and app.Run...
	// All path processed by modules..
	// service.Handle("/", Modules.App)
	registry.Register(service)

	// 通过registry可以获得服务器的ip和端口等信息...
	// find self
	rsp, err := registry.GetService(service.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got service %+v\n", rsp[0])
		fmt.Printf("Nodes info %+v\n", rsp[0].Nodes[0])
	}

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-notify

	registry.Deregister(service)
}
