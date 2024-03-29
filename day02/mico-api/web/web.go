/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-12 10:07:15
 * @LastEditTime: 2019-08-12 10:45:05
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
	)
	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))
	service.HandleFunc("/websocket", hi)
	// websocket interface
	service.HandleFunc("/websocket/hi2", hi2)
	// websocket interface
	service.HandleFunc("/websocket/hi3/hi3", hi2)
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}
func hi2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello! " + r.Form.Get("name"),
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func hi(w http.ResponseWriter, r *http.Request) {

	c, err := upGrader.Upgrade(w, r, nil)
	fmt.Println("进入函数")
	if err != nil {
		fmt.Println("*********")
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
