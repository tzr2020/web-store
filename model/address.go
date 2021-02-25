package model

import "web-store/util"

// Address 用户预存的收货地址表结构
type Address struct {
	ID        int
	UID       int
	Name      string
	Tel       string
	Province  string
	City      string
	Area      string
	Street    string
	Code      string
	IsDefault int // 是否为默认收货地址：0-否，1-是
}

// GetAddressByUserID 从数据库获取用户的所有收货地址，根据用户id
func GetAddressByUserID(uid int) ([]*Address, error) {
	query := "select id, uid, name, tel, province, city, area, street, code, is_default from addresses"
	query += " where uid = ?"

	var addresses []*Address

	rows, err := util.Db.Query(query, uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		address := &Address{}
		err := rows.Scan(&address.ID, &address.UID, &address.Name, &address.Tel,
			&address.Province, &address.City, &address.Area, &address.Street,
			&address.Code, &address.IsDefault)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
