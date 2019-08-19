/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 17:37:00
 * @LastEditTime: 2019-08-18 08:30:41
 * @LastEditors: Please set LastEditors
 */
package basic

import (
	"github.com/yuwe1/micolearn/microservice/basic/config"
	"github.com/yuwe1/micolearn/microservice/basic/db"
	"github.com/yuwe1/micolearn/microservice/basic/redis"
)
func Init(){
	config.Init()
	db.Init()
	redis.Init()
}