/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 15:03:04
 * @LastEditTime: 2019-08-19 15:17:59
 * @LastEditors: Please set LastEditors
 */
package orders

import (
	"context"
	"github.com/yuwe1/micolearn/microservice/basic/common"
	"github.com/yuwe1/micolearn/microservice/basic/db"
	invS  "github.com/yuwe1/micolearn/microservice/inventory-srv/proto/inventory"
	"github.com/micro/go-micro/util/log"
)
// New 新增订单
func (s *service) New(bookId int64, userId int64) (orderId int64, err error) {
	// 请求销存
	rsp, err := invClient.Sell(context.TODO(), &invS.Request{
		BookId: bookId, UserId: userId,
	})
	if err != nil {
		log.Logf("[New] Sell 调用库存服务时失败：%s", err.Error())
		return
	}

	// 获取数据库
	o := db.GetDB()
	insertSQL := `INSERT orders (user_id, book_id, inv_his_id, state) VALUE (?, ?, ?, ?)`

	r, err := o.Exec(insertSQL, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[New] 新增订单失败，err：%s", err)
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

// UpdateOrderState 更新订单状态
func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSQL := `UPDATE orders SET state = ? WHERE id = ?`

	// 获取数据库
	o := db.GetDB()
	// 更新
	_, err = o.Exec(updateSQL, state, orderId)
	if err != nil {
		log.Logf("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}