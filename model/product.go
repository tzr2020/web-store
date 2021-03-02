package model

import (
	"database/sql"
	"errors"
	"log"
	"web-store/util"
)

var (
	// ErrNotFoundProduct 在数据库没有找到产品
	ErrNotFoundProduct = errors.New("在数据库没有找到产品")
)

type Product struct {
	ID          int
	Category_id int
	Name        string
	Price       float64
	Stock       int
	Sales       int
	ImgPath     string
	Detail      string
	HotPoint    string
	Category    *Category // 用于模板
}

// GetProducts 从数据库获取所有产品
func GetProducts() ([]*Product, error) {
	query := "select id, category_id, name, price, stock, sales, img_path, detail, hot_point from products"

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	var ps []*Product

	for rows.Next() {
		p := &Product{}
		rows.Scan(&p.ID, &p.Category_id, &p.Name, &p.Price, &p.Stock, &p.Sales,
			&p.ImgPath, &p.Detail, &p.HotPoint)
		ps = append(ps, p)
	}

	return ps, nil
}

// GetProduct 根据产品id，从数据库获取产品
func GetProduct(pid string) (*Product, error) {
	query := "select id, category_id, name, price, stock, sales, img_path, detail, hot_point from products"
	query += " where id = ?"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		log.Println("准备SQL语句发生错误")
		return nil, err
	}

	p := &Product{}

	err = stmt.QueryRow(pid).Scan(&p.ID, &p.Category_id, &p.Name, &p.Price,
		&p.Stock, &p.Sales, &p.ImgPath, &p.Detail, &p.HotPoint)
	if err == sql.ErrNoRows {
		return nil, ErrNotFoundProduct
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (product *Product) UpdateStockAndSales() error {
	query := "update products set stock = ?, sales = ?"
	query += " where id = ?"

	_, err := util.Db.Exec(query, product.Stock, product.Sales, product.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetProductByID 从数据库获取产品，根据产品id
func GetProductByID(pid int) (*Product, error) {
	query := "select id, category_id, name, price, stock, sales, img_path, detail, hot_point from products"
	query += " where id = ?"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return nil, err
	}

	p := &Product{}

	err = stmt.QueryRow(pid).Scan(&p.ID, &p.Category_id, &p.Name, &p.Price,
		&p.Stock, &p.Sales, &p.ImgPath, &p.Detail, &p.HotPoint)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetIndexNewProducts 从数据库获取首页新品产品
func GetIndexNewProducts() ([]*Product, error) {
	query := "select product_id from index_new_products"

	var products []*Product

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product_id int

		err = rows.Scan(&product_id)
		if err != nil {
			return nil, err
		}

		product, err := GetProductByID(product_id)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetIndexHotProducts 从数据库获取首页热销良品
func GetIndexHotProducts() ([]*Product, error) {
	query := "select product_id from index_hot_products"

	var products []*Product

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product_id int

		err = rows.Scan(&product_id)
		if err != nil {
			return nil, err
		}

		product, err := GetProductByID(product_id)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetIndexRecomProducts 从数据库获取首页推荐产品
func GetIndexRecomProducts() ([]*Product, error) {
	query := "select product_id from index_recom_products"

	var products []*Product

	rows, err := util.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product_id int

		err = rows.Scan(&product_id)
		if err != nil {
			return nil, err
		}

		product, err := GetProductByID(product_id)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
