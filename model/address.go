package model

// Address 用户预存的收货地址表结构
type Address struct {
	ID        int
	UID       int
	Name      string
	Tel       string
	Province  string
	City      string
	Area      string
	Strees    string
	Code      string
	IsDefault int
}
