package model

import (
	"log"
	"web-store/util"
)

// Order 订单结构
type Order struct {
	ID           string
	UID          int     // 用户id
	TotalCount   int     // 订单项数
	TotalAmount  float64 // 产品金额
	Payment      float64 // 支付金额=产品金额+运费
	PaymentType  int     // 支付方式：1-在线支付，2-货到付款
	ShipNumber   string  // 快递单号
	ShipName     string  // 快递公司
	ShipFee      float64 // 运费
	OrderStatus  int     // 状态字典的状态码
	CreateTime   string  // 创建时间
	UpdateTime   string  // 更新时间
	PaymentTime  string  // 支付时间
	ShipTime     string  // 发货时间
	ReceivedTime string  // 收货时间
	FinishTime   string  // 完成时间
	CloseTime    string  // 关闭时间
	Status       int     // 状态：0-禁用，1-正常，-1-删除
}

func (order *Order) Add() error {
	query := "insert into orders (id, uid, total_count, total_amount, payment, payment_type, ship_fee, create_time) values (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		log.Println("准备SQL语句发生错误:", err)
		return err
	}

	_, err = stmt.Exec(order.ID, order.UID, order.TotalCount, order.TotalAmount,
		order.Payment, order.PaymentType, order.ShipFee, order.CreateTime)
	if err != nil {
		return err
	}

	return nil
}

// GetOrdersByUserID 从数据库获取某用户所有订单，根据用户id
func GetOrdersByUserID(uid int) ([]*Order, error) {
	query := "select id, uid, total_count, total_amount, payment, payment_type, ship_number, ship_name, ship_fee, order_status, create_time, update_time, payment_time, ship_time, received_time, finish_time, close_time, status from orders"
	query += " where status = 1 and uid = ?"

	var orders []*Order

	rows, err := util.Db.Query(query, uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.UID, &order.TotalCount, &order.TotalAmount, &order.Payment,
			&order.PaymentType, &order.ShipNumber, &order.ShipName, &order.ShipFee, &order.OrderStatus,
			&order.CreateTime, &order.UpdateTime, &order.PaymentTime, &order.ShipTime, &order.ReceivedTime,
			&order.FinishTime, &order.CloseTime, &order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// UpdateOrderStatus 数据库更新订单的状态字典的状态码
func UpdateOrderStatus(orderID string, toStatus int) (bool, error) {
	query := "update orders set order_status = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, toStatus, orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateOrderPaymentTime 数据库更新订单的付款时间
func UpdateOrderPaymentTime(orderID string, time string) (bool, error) {
	query := "update orders set payment_time = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, time, orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateOrderPaymentTime 数据库更新订单的更新时间
func UpdateOrderUpdateTime(orderID string, time string) (bool, error) {
	query := "update orders set update_time = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, time, orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateOrderPaymentTime 数据库更新订单的付款时间
func UpdateOrderReceivedTime(orderID string, time string) (bool, error) {
	query := "update orders set received_time = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, time, orderID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetOrderByID(orderID string) (*Order, error) {
	query := "select id, uid, total_count, total_amount, payment, payment_type, ship_number, ship_name, ship_fee, order_status, create_time, update_time, payment_time, ship_time, received_time, finish_time, close_time, status from orders"
	query += " where id = ?"

	order := &Order{}

	err := util.Db.QueryRow(query, orderID).Scan(&order.ID, &order.UID, &order.TotalCount,
		&order.TotalAmount, &order.Payment, &order.PaymentType, &order.ShipNumber,
		&order.ShipName, &order.ShipFee, &order.OrderStatus, &order.CreateTime,
		&order.UpdateTime, &order.PaymentTime, &order.ShipTime, &order.ReceivedTime,
		&order.FinishTime, &order.CloseTime, &order.Status)
	if err != nil {
		return nil, err
	}

	return order, nil
}
