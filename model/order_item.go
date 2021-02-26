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
