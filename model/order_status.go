package model

import (
	"log"
	"web-store/util"
)

// OrderStatus 订单状态字典结构
type OrderStatus struct {
	ID   int    `json:"id,string"`
	Code int    `json:"code,string"` // 状态码
	Name string `json:"name"`        // 状态名称
	Text string `json:"text"`        // 状态描述
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

// OrderStatusCodeToOperateURL 是模板函数，用于将订单状态字典的状态代码转换为对应的操作URL
func OrderStatusCodeToOperateURL(code int) string {

	switch code {
	case 0, 1, 2: // 等待买家付款
		return "/payOrder"
	case 4: // 卖家已发货，等待买家确认收货
		return "/receivedOrder"
	}

	return "#"
}

// OrderStatusCodeToOperateText 是模板函数，用于将订单状态字典的状态代码转换为对应的操作文本
func OrderStatusCodeToOperateText(code int) string {

	switch code {
	case 0, 1, 2: // 等待买家付款
		return "去付款"
	case 4: // 卖家已发货，等待买家确认收货
		return "确认收货"
	}

	return "暂无操作"
}

func GetOrderStatusByCode(code int) (*OrderStatus, error) {
	query := `select id, code, name, text
		from order_status
		where code=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return nil, err
	}

	osobj := &OrderStatus{}

	err = stmt.QueryRow(code).Scan(&osobj.ID, &osobj.Code, &osobj.Name, &osobj.Text)
	if err != nil {
		return nil, err
	}

	return osobj, nil
}

func GetOrderStatusPage(pageNo int, pageSize int) (osobjs []*OrderStatus, err error) {
	query := `select id, code, name, text
		from order_status
		limit ?,?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		osobj := &OrderStatus{}
		err := rows.Scan(&osobj.ID, &osobj.Code, &osobj.Name, &osobj.Text)
		if err != nil {
			return nil, err
		}
		osobjs = append(osobjs, osobj)
	}

	return
}

func (osobj *OrderStatus) Add() error {
	query := `insert into order_status (code, name, text)
		values (?,?,?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(osobj.Code, osobj.Name, osobj.Text)
	if err != nil {
		return err
	}

	return nil
}

func (osobj *OrderStatus) Update() error {
	query := `update order_status set code=?, name=?, text=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(osobj.Code, osobj.Name, osobj.Text, osobj.ID)
	if err != nil {
		return err
	}

	return nil
}

func (osobj *OrderStatus) Delete() error {
	query := `delete from order_status where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(osobj.ID)
	if err != nil {
		return err
	}

	return nil
}
