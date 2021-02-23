package model

import "web-store/util"

type Category struct {
	ID    int
	Name  string // 产品类别名称
	PID   int    // 父级id
	Level int    // 层级
	Img   string // 图片路径
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
