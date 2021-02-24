package model

// Order 订单结构
type Order struct {
	ID           string
	UID          int         // 用户id
	TotalCount   int         // 订单项数
	TotalAmount  float64     // 产品金额
	Payment      float64     // 支付金额=产品金额+运费
	PaymentType  int         // 支付方式：1-在线支付，2-货到付款
	ShipNumber   int         // 快递单号
	ShipName     string      // 快递公司
	ShipFee      float64     // 运费
	OrderStatus  OrderStatus // 状态字典
	CreateTime   string      // 创建时间
	UpdateTime   string      // 更新时间
	PaymentTime  string      // 支付时间
	ShipTime     string      // 发货时间
	ReceivedTime string      // 收货时间
	FinishTime   string      // 完成时间
	CloseTime    string      // 关闭时间
	Status       int         // 状态：0-禁用，1-正常，-1-删除
}
