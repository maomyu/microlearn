/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 17:37:00
 * @LastEditTime: 2019-08-17 08:46:55
 * @LastEditors: Please set LastEditors
 */
package basic

import (
	"github.com/yuwe1/micolearn/microservice/user-srv/basic/config"
	"github.com/yuwe1/micolearn/microservice/user-srv/basic/db"
)
func Init(){
	config.Init()
	db.Init()
}