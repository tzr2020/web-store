package model

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"web-store/util"
)

var (
	// ErrNotFoundProduct 在数据库没有找到产品
	ErrNotFoundProduct = errors.New("在数据库没有找到产品")
)

type Product struct {
	ID          int       `json:"id,string"`
	Category_id int       `json:"categoryID,string"`
	Name        string    `json:"name"`
	Price       float64   `json:"price,string"`
	Stock       int       `json:"stock,string"`
	Sales       int       `json:"sales,string"`
	ImgPath     string    `json:"imgPath"`
	Detail      string    `json:"detail"`
	HotPoint    string    `json:"hotPoint"`
	Category    *Category `json:"category"` // 用于模板
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

// GetNavProducts 从数据库获取导航栏产品
func GetNavProducts() ([]*Product, error) {
	query := "select product_id from nav_products"

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

// GetProductPage 查询数据库，获取产品列表，根据当前页的页码和每页记录条数
func GetProductPage(pageNo int, pageSize int) ([]*Product, error) {
	query := `select id, category_id, name, price, stock, sales, img_path, 
		detail, hot_point from products limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var ps []*Product
	for rows.Next() {
		p := &Product{}
		err = rows.Scan(&p.ID, &p.Category_id, &p.Name, &p.Price,
			&p.Stock, &p.Sales, &p.ImgPath, &p.Detail, &p.HotPoint)
		if err != nil {
			return nil, err
		}

		categoryID := strconv.Itoa(p.Category_id)
		cate, err := GetCategory(categoryID)
		if err != nil {
			return nil, err
		}
		p.Category = cate

		ps = append(ps, p)
	}

	return ps, nil
}

// Add 数据库添加产品
func (p Product) Add() error {
	query := `insert into products (category_id, name, price, stock, sales, hot_point)
		values (?,?,?,?,?,?)`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&p.Category_id, &p.Name, &p.Price, &p.Stock, &p.Sales, &p.HotPoint)
	if err != nil {
		return err
	}

	return nil
}

// Update 数据库更新产品
func (p Product) Update() error {
	query := `update products set category_id=?, name=?, price=?, stock=?, sales=?, hot_point=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&p.Category_id, &p.Name, &p.Price, &p.Stock, &p.Sales, &p.HotPoint, &p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete 数据库删除产品
func (p Product) Delete() error {
	query := `delete from products where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProductImg 数据库更新产品图片
func UpdateProductImg(imgPath string, uid string) (err error) {
	query := `update products set img_path = ? where id = ?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(imgPath, uid)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProductDetail 数据库更新产品详情
func UpdateProductDetail(detail string, uid string) (err error) {
	query := `update products set detail = ? where id = ?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(detail, uid)
	if err != nil {
		return err
	}

	return nil
}
