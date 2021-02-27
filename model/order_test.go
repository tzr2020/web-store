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
	// t.Run("测试从数据库获取所有订单支付方式", testGetOrderPaymentType)
	// t.Run("测试从数据库获取订单的状态字典", testGetOrderStatus)
	// t.Run("测试从数据库获取某用户的所有订单，根据用户id", testGetOrdersByUserID)
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

func testGetOrderPaymentType(t *testing.T) {
	fmt.Println("测试从数据库获取所有订单支付方式")

	orderPaymentTypes, err := GetOrderPaymentTypes()
	if err != nil {
		log.Println("从数据库获取所有订单支付方式发生错误")
	}

	for _, v := range orderPaymentTypes {
		fmt.Println(v)
	}
}

func testGetOrderStatus(t *testing.T) {
	fmt.Println("测试从数据库获取订单的状态字典")

	allStatus, err := GetOrderStatus()
	if err != nil {
		log.Println("从数据库获取订单的状态字典发生错误")
	}

	for _, v := range allStatus {
		fmt.Println(v)
	}
}

func testGetOrdersByUserID(t *testing.T) {
	fmt.Println("测试从数据库获取某用户的所有订单，根据用户id")

	orders, err := GetOrdersByUserID(1)
	if err != nil {
		log.Println("从数据库获取某用户的所有订单发生错误:", err)
		return
	}

	for _, v := range orders {
		fmt.Println(v)
	}
}
