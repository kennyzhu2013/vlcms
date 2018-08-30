/*
@Time : 2018/8/30 15:56 
@Author : kenny zhu
@File : main.go
@Software: GoLand
@Others:
*/
package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/pborman/uuid"
	"strconv"
	"github.com/kataras/iris"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	service = &registry.Service{
		Name: "go.micro.srv.iris",
		Nodes: []*registry.Node{
			{
				Id:      "go.micro.srv.iris-" + uuid.NewUUID().String(),
				Address: "localhost",
				Port:    8400,
			},
		},
	}
)

// This is example for api gates..
// To delete....
func main() {
	conf.Init( "./conf/irisapi.json" )
	service.Name = conf.AppConf.ApiName

	// register call back.
	Modules.App.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	})

	address := conf.AppConf.IpAddress + ":" + strconv.Itoa( int(conf.AppConf.Port) )
	go Start( iris.Addr( address ) ) // eg:":8400"
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

