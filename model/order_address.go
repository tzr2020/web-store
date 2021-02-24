package model

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
