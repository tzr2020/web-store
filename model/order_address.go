package model

import (
	"log"
	"web-store/util"
)

// OrderAddress 订单地址结构
type OrderAddress struct {
	ID       int    `json:"id,string"`
	OrderID  string `json:"order_id"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Street   string `json:"street"`
	Code     string `json:"code"`
	Address  string `json:"address"`
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
		orderAddress.Street, orderAddress.Code)
	if err != nil {
		return err
	}

	return nil
}

// GetOrderAddressByOrderID 从数据库获取订单地址，根据订单id
func GetOrderAddressByOrderID(orderID string) (*OrderAddress, error) {
	query := "select id, order_id, name, tel, province, city, area, street, code from order_addresses"
	query += " where order_id = ?"

	orderAddress := &OrderAddress{}

	err := util.Db.QueryRow(query, orderID).Scan(&orderAddress.ID, &orderAddress.OrderID,
		&orderAddress.Name, &orderAddress.Tel, &orderAddress.Province, &orderAddress.City,
		&orderAddress.Area, &orderAddress.Street, &orderAddress.Code)
	if err != nil {
		return nil, err
	}

	return orderAddress, nil
}

func GetOrderAddressPage(pageNo int, pageSize int) ([]*OrderAddress, error) {
	query := `select id, order_id, name, tel, province, city, area, street, code
		from order_addresses
		limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var oas []*OrderAddress
	for rows.Next() {
		oa := &OrderAddress{}
		err = rows.Scan(&oa.ID, &oa.OrderID, &oa.Name, &oa.Tel, &oa.Province,
			&oa.City, &oa.Area, &oa.Street, &oa.Code)
		if err != nil {
			return nil, err
		}

		oas = append(oas, oa)
	}

	return oas, nil
}

func (oa OrderAddress) Update() error {
	query := `update order_addresses 
		set order_id=?, name=?, tel=?, province=?, city=?, area=?, street=?, code=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&oa.OrderID, &oa.Name, &oa.Tel, &oa.Province,
		&oa.City, &oa.Area, &oa.Street, &oa.Code, &oa.ID)
	if err != nil {
		return err
	}

	return nil
}

func (oa OrderAddress) Delete() error {
	query := `delete from order_addresses 
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&oa.ID)
	if err != nil {
		return err
	}

	return nil
}
