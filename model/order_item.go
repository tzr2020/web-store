package model

import (
	"log"
	"web-store/util"
)

// OrderItem 订单项结构
type OrderItem struct {
	ID        int      `json:"id,string"`
	OrderID   string   `json:"order_id"`
	Order     *Order   `json:"order"`
	ProductID int      `json:"product_id,string"`
	Product   *Product `json:"product"`
	Count     int      `json:"count,string"`
	Amount    float64  `json:"amount,string"`
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

func GetOrderItemPage(pageNo int, pageSize int) ([]*OrderItem, error) {
	query := `select id, order_id, product_id, count, amount
		from order_items
		limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var oits []*OrderItem
	for rows.Next() {
		oit := &OrderItem{}
		err = rows.Scan(&oit.ID, &oit.OrderID, &oit.ProductID,
			&oit.Count, &oit.Amount)
		if err != nil {
			return nil, err
		}

		o, err := GetOrderByID(oit.OrderID)
		if err != nil {
			return nil, err
		}
		oit.Order = o

		p, err := GetProductByID(oit.ProductID)
		if err != nil {
			return nil, err
		}
		oit.Product = p

		oits = append(oits, oit)
	}

	return oits, nil
}

func (oit OrderItem) Update() error {
	query := `update order_items 
		set order_id=?, product_id=?, count=?, amount=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&oit.OrderID, &oit.ProductID,
		&oit.Count, &oit.Amount, &oit.ID)
	if err != nil {
		return err
	}

	return nil
}

func (oit OrderItem) Delete() error {
	query := `delete from order_items
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&oit.ID)
	if err != nil {
		return err
	}

	return nil
}
