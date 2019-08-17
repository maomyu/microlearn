/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 09:39:23
 * @LastEditTime: 2019-08-17 09:55:48
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/user-srv/basic"
	"github.com/yuwe1/micolearn/microservice/user-srv/basic/config"
	"github.com/yuwe1/micolearn/microservice/user-srv/handler"
	"github.com/yuwe1/micolearn/microservice/user-srv/model"
	s "github.com/yuwe1/micolearn/microservice/user-srv/proto/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 初始化配置
	basic.Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型
			model.Init()
			handler.Init()
		}),
	)
	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
