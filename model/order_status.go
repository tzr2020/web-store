package model

// OrderStatus 订单状态字典结构
type OrderStatus struct {
	ID   int
	Code int    // 状态码
	Name string // 状态名称
	Text string // 状态描述
}
