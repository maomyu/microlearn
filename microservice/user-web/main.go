/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 10:50:36
 * @LastEditTime: 2019-08-17 11:01:06
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/yuwe1/micolearn/microservice/user-web/basic"
	"github.com/yuwe1/micolearn/microservice/user-web/basic/config"
	"github.com/yuwe1/micolearn/microservice/user-web/handler"
)

func main() {
	// 初始化配置
	basic.Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)
	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
