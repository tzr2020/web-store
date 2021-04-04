package model

import (
	"strconv"
	"web-store/util"
)

type IndexNewProduct struct {
	ID        int     `json:"id,omitempty,string"`
	ProductID int     `json:"product_id,omitempty,string"`
	Product   Product `json:"product,omitempty"`
}

// GetIndexNewProductPage 查询数据库，获取首页最新产品列表，根据当前页的页码和每页记录条数
func GetIndexNewProductPage(pageNo int, pageSize int) (inps []*IndexNewProduct, err error) {
	query := `select id, product_id
		from index_new_products limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		inp := &IndexNewProduct{}
		err = rows.Scan(&inp.ID, &inp.ProductID)
		if err != nil {
			return nil, err
		}

		p, err := GetProductByID(inp.ProductID)
		if err != nil {
			return nil, err
		}

		categoryID := strconv.Itoa(p.Category_id)
		cate, err := GetCategory(categoryID)
		if err != nil {
			return nil, err
		}
		p.Category = cate

		inp.Product = *p
		inps = append(inps, inp)
	}

	return
}

// Add 数据库添加首页最新产品
func (inp IndexNewProduct) Add() error {
	query := `insert into index_new_products (product_id)
		values (?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(inp.ProductID)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新首页最新产品
func (inp IndexNewProduct) Update() error {
	query := `update index_new_products set product_id=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(inp.ProductID, inp.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除首页最新产品
func (inp IndexNewProduct) Delete() error {
	query := `delete from index_new_products where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(inp.ID)
	if err != nil {
		return err
	}

	return nil
}
