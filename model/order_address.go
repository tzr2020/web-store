package model

import (
	"log"
	"web-store/util"
)

// OrderAddress 订单地址结构
type OrderAddress struct {
	ID       int
	OrderID  string
	Name     string
	Tel      string
	Province string
	City     string
	Area     string
	Strees   string
	Code     string
}

func (orderAddress *OrderAddress) Add() error {
	query := "insert into order_addresses (order_id, name, tel, province, city, area, street, code) values (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		log.Println("准备SQL语句发生错误:", err)
		return err
	}

	_, err = stmt.Exec(orderAddress.OrderID, orderAddress.Name, orderAddress.Tel,
		orderAddress.Province, orderAddress.City, orderAddress.Area,
		orderAddress.Strees, orderAddress.Code)
	if err != nil {
		return err
	}

	return nil
}
