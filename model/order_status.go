package model

import (
	"log"
	"web-store/util"
)

// OrderStatus 订单状态字典结构
type OrderStatus struct {
	ID   int
	Code int    // 状态码
	Name string // 状态名称
	Text string // 状态描述
}

// GetOrderStatus 从数据库获取订单的状态字典
func GetOrderStatus() ([]*OrderStatus, error) {
	query := "select id, code, name, text from order_status"

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	var allOrderStatus []*OrderStatus

	for rows.Next() {
		orderStatus := &OrderStatus{}
		rows.Scan(&orderStatus.ID, &orderStatus.Code,
			&orderStatus.Name, &orderStatus.Text)
		if err != nil {
			return nil, err
		}
		allOrderStatus = append(allOrderStatus, orderStatus)
	}

	return allOrderStatus, nil
}

// OrderStatusCodeToText 是模板函数，用于将订单状态字典的状态代码转换为对应的状态文本
func OrderStatusCodeToText(code int) string {
	allStatus, err := GetOrderStatus()
	if err != nil {
		log.Println("从数据库获取订单状态字典发生错误:", err)
		return ""
	}

	for _, v := range allStatus {
		if code == v.Code {
			return v.Text
		}
	}

	log.Println("数据库没有匹配的订单状态字典的状态代码")
	return ""
}
