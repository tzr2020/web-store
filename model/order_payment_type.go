package model

import (
	"log"
	"web-store/util"
)

type OrderPaymentType struct {
	ID   int    `json:"id,string"`
	Code int    `json:"code,string"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// GetOrderPaymentTypes 从数据库获取所有订单支付方式
func GetOrderPaymentTypes() ([]*OrderPaymentType, error) {
	query := "select id, code, name, text from order_payment_type"

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	var orderPaymentTypes []*OrderPaymentType

	for rows.Next() {
		orderPaymentType := &OrderPaymentType{}
		rows.Scan(&orderPaymentType.ID, &orderPaymentType.Code,
			&orderPaymentType.Name, &orderPaymentType.Text)
		if err != nil {
			return nil, err
		}
		orderPaymentTypes = append(orderPaymentTypes, orderPaymentType)
	}

	return orderPaymentTypes, nil
}

// OrderPaymentTypeCodeToText 是模板函数，用于将订单的支付方式代码转换为对应的支付方式文本
func OrderPaymentTypeCodeToText(code int) string {

	types, err := GetOrderPaymentTypes()
	if err != nil {
		log.Println("从数据库获取所有订单的支付方式发生错误:", err)
		return ""
	}

	for _, v := range types {
		if code == v.Code {
			return v.Text
		}
	}

	log.Println("数据库没有匹配的支付方式代码")
	return ""
}

func GetOrderPaymentTypeByCode(code int) (*OrderPaymentType, error) {
	query := `select id, code, name, text 
		from order_payment_type
		where code=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return nil, err
	}

	opt := &OrderPaymentType{}

	err = stmt.QueryRow(code).Scan(&opt.ID, &opt.Code, &opt.Name, &opt.Text)
	if err != nil {
		return nil, err
	}

	return opt, nil
}

func GetOrderPaymentTypePage(pageNo int, pageSize int) (opts []*OrderPaymentType, err error) {
	query := `select id, code, name, text
		from order_payment_type
		limit ?,?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		opt := &OrderPaymentType{}
		err := rows.Scan(&opt.ID, &opt.Code, &opt.Name, &opt.Text)
		if err != nil {
			return nil, err
		}
		opts = append(opts, opt)
	}

	return
}

func (opt OrderPaymentType) Add() error {
	query := `insert into order_payment_type (code, name, text)
		values (?,?,?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(opt.Code, opt.Name, opt.Text)
	if err != nil {
		return err
	}

	return nil
}

func (opt *OrderPaymentType) Update() error {
	query := `update order_payment_type set code=?, name=?, text=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(opt.Code, opt.Name, opt.Text, opt.ID)
	if err != nil {
		return err
	}

	return nil
}

func (opt *OrderPaymentType) Delete() error {
	query := `delete from order_payment_type where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(opt.ID)
	if err != nil {
		return err
	}

	return nil
}
