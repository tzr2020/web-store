package model

import "web-store/util"

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
}

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
