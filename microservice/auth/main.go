/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-18 09:04:15
 * @LastEditTime: 2019-08-18 12:29:21
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	"github.com/yuwe1/micolearn/microservice/auth/handler"
	"github.com/yuwe1/micolearn/microservice/auth/model"
	s "github.com/yuwe1/micolearn/microservice/auth/proto/auth"
	"github.com/yuwe1/micolearn/microservice/basic"
	"github.com/yuwe1/micolearn/microservice/basic/config"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
)

func main() {
	//  初始化配置
	basic.Init()
	// 使用conssul进行注册
	micReg := consul.NewRegistry(registryOptions)
	// 创建新服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)
	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)
	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
