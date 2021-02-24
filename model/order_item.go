package model

// OrderItem 订单项结构
type OrderItem struct {
	ID        int
	OrderID   string
	ProductID *Product
	Count     int
	Amount    float64
}
