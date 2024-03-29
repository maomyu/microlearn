/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 16:18:12
 * @LastEditTime: 2019-08-19 17:43:23
 * @LastEditors: Please set LastEditors
 */
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	auth "github.com/yuwe1/micolearn/microservice/auth/proto/auth"
	invS "github.com/yuwe1/micolearn/microservice/inventory-srv/proto/inventory"
	orders "github.com/yuwe1/micolearn/microservice/orders-srv/proto/orders"
	"github.com/yuwe1/micolearn/microservice/plugins/session"
)

var (
	serviceClient orders.OrdersService
	authClient    auth.Service
	invClient     invS.InventoryService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = orders.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// 新增订单入口
func New(w http.ResponseWriter, r *http.Request) {
	// 仅仅接受post请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)
	fmt.Println("bookID:", bookId)
	// 返回的结果
	response := map[string]interface{}{}
	// 开始调用后台服务
	fmt.Println("userID", session.GetSession(w, r).Values["userId"].(int64))
	rsp, err := serviceClient.New(context.TODO(), &orders.Request{
		BookId: bookId,
		UserId: session.GetSession(w, r).Values["userId"].(int64),
	})
	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

//
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
