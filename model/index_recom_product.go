package model

import (
	"strconv"
	"web-store/util"
)

type IndexRecomProduct struct {
	ID        int     `json:"id,omitempty,string"`
	ProductID int     `json:"product_id,omitempty,string"`
	Product   Product `json:"product,omitempty"`
}

// GetIndexRecomProductPage 查询数据库，获取首页推荐产品列表，根据当前页的页码和每页记录条数
func GetIndexRecomProductPage(pageNo int, pageSize int) (irps []*IndexRecomProduct, err error) {
	query := `select id, product_id
		from index_recom_products limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		irp := &IndexRecomProduct{}
		err = rows.Scan(&irp.ID, &irp.ProductID)
		if err != nil {
			return nil, err
		}

		p, err := GetProductByID(irp.ProductID)
		if err != nil {
			return nil, err
		}

		categoryID := strconv.Itoa(p.Category_id)
		cate, err := GetCategory(categoryID)
		if err != nil {
			return nil, err
		}
		p.Category = cate

		irp.Product = *p
		irps = append(irps, irp)
	}

	return
}

// Add 数据库添加首页推荐产品
func (irp IndexRecomProduct) Add() error {
	query := `insert into index_recom_products (product_id)
		values (?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(irp.ProductID)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新首页推荐产品
func (irp IndexRecomProduct) Update() error {
	query := `update index_recom_products set product_id=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(irp.ProductID, irp.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除首页推荐产品
func (irp IndexRecomProduct) Delete() error {
	query := `delete from index_recom_products where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(irp.ID)
	if err != nil {
		return err
	}

	return nil
}
