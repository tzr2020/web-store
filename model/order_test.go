package model

import (
	"fmt"
	"log"
	"testing"
	"time"
	"web-store/util"
)

func TestOrder(t *testing.T) {
	// t.Run("测试添加订单", testAddOrder)
}

func testAddOrder(t *testing.T) {
	fmt.Println("测试添加订单")

	orderID := util.CreateUUID()

	oit := &OrderItem{
		OrderID:   orderID,
		ProductID: 1,
		Count:     1,
		Amount:    30.5,
	}
	oit2 := &OrderItem{
		OrderID:   orderID,
		ProductID: 2,
		Count:     1,
		Amount:    20.0,
	}

	order := &Order{
		ID:          orderID,
		UID:         1,
		TotalCount:  2,
		TotalAmount: 50.5,
		Payment:     61.5,
		PaymentType: 2,
		ShipFee:     11.0,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	orderAddress := &OrderAddress{
		OrderID:  orderID,
		Name:     "张三",
		Tel:      "13588746367",
		Province: "山东",
		City:     "青岛",
		Area:     "城阳区",
		Strees:   "xxx路xx号",
		Code:     "345345",
	}

	if err := order.Add(); err != nil {
		log.Println("数据库发生错误:", err)
		return
	}
	if err := oit.Add(); err != nil {
		log.Println("数据库发生错误:", err)
		return
	}
	if err := oit2.Add(); err != nil {
		log.Println("数据库发生错误:", err)
		return
	}
	if err := orderAddress.Add(); err != nil {
		log.Println("数据库发生错误:", err)
		return
	}
}
