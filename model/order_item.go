package model

import (
	"log"
	"web-store/util"
)

// OrderItem 订单项结构
type OrderItem struct {
	ID        int
	OrderID   string
	ProductID int
	Count     int
	Amount    float64
	Product   *Product
}

func (oit *OrderItem) Add() error {
	query := "insert into order_items (order_id, Product_id, count, amount) values (?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		log.Println("准备SQL语句发生错误:", err)
		return err
	}

	_, err = stmt.Exec(oit.OrderID, oit.ProductID, oit.Count, oit.Amount)
	if err != nil {
		return err
	}

	return nil
}

// GetOrderItemsByOrderID 从数据库获取订单项，根据订单id
func GetOrderItemsByOrderID(orderID string) ([]*OrderItem, error) {
	query := "select id, order_id, Product_id, count, amount from order_items"
	query += " where order_id = ?"

	var orderItems []*OrderItem

	rows, err := util.Db.Query(query, orderID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		orderItem := &OrderItem{}
		rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID,
			&orderItem.Count, &orderItem.Amount)
		if err != nil {
			return nil, err
		}

		// 将从数据库获取的产品设置到订单项结构的产品字段
		product, err := GetProductByID(orderItem.ProductID)
		if err != nil {
			return nil, err
		}
		orderItem.Product = product

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
