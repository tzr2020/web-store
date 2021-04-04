package model

import (
	"strconv"
	"web-store/util"
)

type NavProduct struct {
	ID        int     `json:"id,omitempty,string"`
	ProductID int     `json:"product_id,omitempty,string"`
	Product   Product `json:"product,omitempty"`
}

// GetNavProductPage 查询数据库，获取导航栏产品列表，根据当前页的页码和每页记录条数
func GetNavProductPage(pageNo int, pageSize int) (nps []*NavProduct, err error) {
	query := `select id, product_id
		from nav_products limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		np := &NavProduct{}
		err = rows.Scan(&np.ID, &np.ProductID)
		if err != nil {
			return nil, err
		}

		p, err := GetProductByID(np.ProductID)
		if err != nil {
			return nil, err
		}

		categoryID := strconv.Itoa(p.Category_id)
		cate, err := GetCategory(categoryID)
		if err != nil {
			return nil, err
		}
		p.Category = cate

		np.Product = *p
		nps = append(nps, np)
	}

	return
}

// Add 数据库添加导航栏产品
func (np NavProduct) Add() error {
	query := `insert into nav_products (product_id)
		values (?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(np.ProductID)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新导航栏产品
func (np NavProduct) Update() error {
	query := `update nav_products set product_id=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(np.ProductID, np.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除导航栏产品
func (np NavProduct) Delete() error {
	query := `delete from nav_products where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(np.ID)
	if err != nil {
		return err
	}

	return nil
}
