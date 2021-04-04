package model

import "web-store/util"

type Category struct {
	ID    int    `json:"id,string"`
	Name  string `json:"name"`         // 产品类别名称
	PID   int    `json:"pid,string"`   // 父级id
	Level int    `json:"level,string"` // 层级
	Img   string `json:"img"`          // 图片路径
}

// GetCategories 从数据库获取所有产品类别
func GetCategories() ([]*Category, error) {
	query := "select id, name, p_id, level, img from categories"

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	var cates []*Category

	for rows.Next() {
		cate := &Category{}
		rows.Scan(&cate.ID, &cate.Name, &cate.PID, &cate.Level, &cate.Img)
		cates = append(cates, cate)
	}

	return cates, nil
}

// GetCategory 根据产品类别id，从数据库获取产品类别
func GetCategory(cate_id string) (*Category, error) {
	query := "select id, name, p_id, level, img from categories"
	query += " where id = ?"

	cate := &Category{}

	err := util.Db.QueryRow(query, cate_id).Scan(&cate.ID, &cate.Name,
		&cate.PID, &cate.Level, &cate.Img)
	if err != nil {
		return nil, err
	}

	return cate, nil
}

// Add 数据库添加产品类别
func (cate Category) Add() error {
	query := `insert into categories (id, name, p_id, level)
		values (?,?,?,?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cate.ID, cate.Name, cate.PID, cate.Level)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新产品类别
func (cate Category) Update() error {
	query := `update categories set name=?, p_id=?, level=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cate.Name, cate.PID, cate.Level, cate.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除产品类别
func (cate Category) Delete() error {
	query := `delete from categories where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cate.ID)
	if err != nil {
		return err
	}

	return nil
}
