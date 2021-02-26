package model

import "web-store/util"

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
