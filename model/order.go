package model

import (
	"time"
	"web-store/util"
)

// Order 订单结构
type Order struct {
	ID               string            `json:"id"`
	UID              int               `json:"uid,string"`          // 用户id
	TotalCount       int               `json:"total_count,string"`  // 订单项数
	TotalAmount      float64           `json:"total_amount,string"` // 产品金额
	Payment          float64           `json:"payment,string"`      // 支付金额=产品金额+运费
	PaymentType      int               `json:"payment_type,string"` // 支付方式
	OrderPaymentType *OrderPaymentType `json:"order_payment_type"`
	ShipNumber       string            `json:"ship_number"`         // 快递单号
	ShipName         string            `json:"ship_name"`           // 快递公司
	ShipFee          float64           `json:"ship_fee,string"`     // 运费
	OrderStatus      int               `json:"order_status,string"` // 状态字典的状态码
	OrderStatusObj   *OrderStatus      `json:"order_status_obj"`
	CreateTime       string            `json:"create_time"`   // 创建时间
	UpdateTime       string            `json:"update_time"`   // 更新时间
	PaymentTime      string            `json:"payment_time"`  // 支付时间
	ShipTime         string            `json:"ship_time"`     // 发货时间
	ReceivedTime     string            `json:"received_time"` // 收货时间
	FinishTime       string            `json:"finish_time"`   // 完成时间
	CloseTime        string            `json:"close_time"`    // 关闭时间
	Status           int               `json:"status,string"` // 状态：0-禁用，1-正常，-1-删除
}

func (order *Order) Add() error {
	query := "insert into orders (id, uid, total_count, total_amount, payment, payment_type, ship_fee, create_time) values (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
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

func GetOrderPage(pageNo int, pageSize int) (os []*Order, err error) {
	query := `select id, uid, total_count, total_amount, payment, payment_type, 
		ship_number, ship_name, ship_fee, order_status, create_time, update_time, 
		payment_time, ship_time, received_time, finish_time, close_time, status 
		from orders
		limit ?,?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.UID, &order.TotalCount, &order.TotalAmount,
			&order.Payment, &order.PaymentType, &order.ShipNumber, &order.ShipName,
			&order.ShipFee, &order.OrderStatus, &order.CreateTime, &order.UpdateTime,
			&order.PaymentTime, &order.ShipTime, &order.ReceivedTime, &order.FinishTime,
			&order.CloseTime, &order.Status)
		if err != nil {
			return nil, err
		}

		opt, err := GetOrderPaymentTypeByCode(order.PaymentType)
		if err != nil {
			return nil, err
		}
		order.OrderPaymentType = opt

		osobj, err := GetOrderStatusByCode(order.OrderStatus)
		if err != nil {
			return nil, err
		}
		order.OrderStatusObj = osobj

		os = append(os, order)
	}

	return
}

func (order *Order) Add2() error {
	query := `insert into orders
		values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	order.ID = util.CreateUUID()
	order.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	order.Payment = order.TotalAmount + order.ShipFee

	_, err = stmt.Exec(&order.ID, &order.UID, &order.TotalCount, &order.TotalAmount,
		&order.Payment, &order.PaymentType, &order.ShipNumber, &order.ShipName,
		&order.ShipFee, &order.OrderStatus, &order.CreateTime, &order.UpdateTime,
		&order.PaymentTime, &order.ShipTime, &order.ReceivedTime, &order.FinishTime,
		&order.CloseTime, &order.Status)
	if err != nil {
		return err
	}

	return nil
}

func (order Order) Update() error {
	query := `update orders set uid=?, total_count=?, total_amount=?,
		payment=?, payment_type=?, ship_number=?, ship_name=?, ship_fee=?, 
		order_status=?, create_time=?, update_time=?, payment_time=?,
		ship_time=?, received_time=?, finish_time=?, close_time=?, status=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	order.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	order.Payment = order.TotalAmount + order.ShipFee

	_, err = stmt.Exec(&order.UID, &order.TotalCount, &order.TotalAmount,
		&order.Payment, &order.PaymentType, &order.ShipNumber, &order.ShipName,
		&order.ShipFee, &order.OrderStatus, &order.CreateTime, &order.UpdateTime,
		&order.PaymentTime, &order.ShipTime, &order.ReceivedTime, &order.FinishTime,
		&order.CloseTime, &order.Status, &order.ID)
	if err != nil {
		return err
	}

	return nil
}

func (order *Order) Delete() error {
	query := `update orders set status=-1
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID)
	if err != nil {
		return err
	}

	return nil
}
