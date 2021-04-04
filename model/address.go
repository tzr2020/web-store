package model

import "web-store/util"

// Address 用户预存的收货地址表结构
type Address struct {
	ID        int    `json:"id,string"`
	UID       int    `json:"uid,string"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Area      string `json:"area"`
	Street    string `json:"street"`
	Code      string `json:"code"`
	IsDefault int    `json:"is_default,string"` // 是否为默认收货地址：0-否，1-是
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

// GetProductPage 查询数据库，获取用户地址列表，根据当前页的页码和每页记录条数
func GetAddressPage(pageNo int, pageSize int) (as []*Address, err error) {
	query := `select id, uid, name, tel, province, city, area, street, code, is_default
		from addresses limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := &Address{}
		err = rows.Scan(&a.ID, &a.UID, &a.Name, &a.Tel, &a.Province,
			&a.City, &a.Area, &a.Street, &a.Code, &a.IsDefault)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}

	return
}

// Add 数据库添加用户地址
func (a Address) Add() (err error) {
	query := `insert into addresses (uid, name, tel, province, city, area, street,
		code, is_default) values (?,?,?,?,?,?,?,?,?)`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(&a.UID, &a.Name, &a.Tel, &a.Province, &a.City, &a.Area, &a.Street,
		&a.Code, &a.IsDefault)
	if err != nil {
		return
	}

	return
}

// Update 数据库更新用户地址
func (a Address) Update() (err error) {
	query := `update addresses set name=?, tel=?, province=?, city=?, area=?, street=?, code=?, is_default=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(&a.Name, &a.Tel, &a.Province, &a.City, &a.Area, &a.Street,
		&a.Code, &a.IsDefault, &a.ID)
	if err != nil {
		return
	}

	return
}

// Update 数据库更新用户地址
func (a Address) Delete() (err error) {
	query := `delete from addresses where id=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(&a.ID)
	if err != nil {
		return
	}

	return
}
