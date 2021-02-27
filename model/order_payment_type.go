package model

import (
	"log"
	"web-store/util"
)

type OrderPaymentType struct {
	ID   int
	Code int
	Name string
	Text string
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
