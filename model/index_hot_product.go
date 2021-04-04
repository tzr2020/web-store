package model

import (
	"strconv"
	"web-store/util"
)

type IndexHotProduct struct {
	ID        int     `json:"id,omitempty,string"`
	ProductID int     `json:"product_id,omitempty,string"`
	Product   Product `json:"product,omitempty"`
}

// GetIndexHotProductPage 查询数据库，获取首页热卖产品列表，根据当前页的页码和每页记录条数
func GetIndexHotProductPage(pageNo int, pageSize int) (ihps []*IndexHotProduct, err error) {
	query := `select id, product_id
		from index_hot_products limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ihp := &IndexHotProduct{}
		err = rows.Scan(&ihp.ID, &ihp.ProductID)
		if err != nil {
			return nil, err
		}

		p, err := GetProductByID(ihp.ProductID)
		if err != nil {
			return nil, err
		}

		categoryID := strconv.Itoa(p.Category_id)
		cate, err := GetCategory(categoryID)
		if err != nil {
			return nil, err
		}
		p.Category = cate

		ihp.Product = *p
		ihps = append(ihps, ihp)
	}

	return
}

// Add 数据库添加首页热卖产品
func (ihp IndexHotProduct) Add() error {
	query := `insert into index_hot_products (product_id)
		values (?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ihp.ProductID)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新首页热卖产品
func (ihp IndexHotProduct) Update() error {
	query := `update index_hot_products set product_id=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ihp.ProductID, ihp.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除首页热卖产品
func (ihp IndexHotProduct) Delete() error {
	query := `delete from index_hot_products where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ihp.ID)
	if err != nil {
		return err
	}

	return nil
}
