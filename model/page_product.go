package model

import (
	"strconv"
	"web-store/util"
)

type PageProduct struct {
	Products    []*Product // 当前页产品列表
	PageNo      int64      // 当前页页码
	PageSize    int64      // 每页记录数
	TotalPageNo int64      // 总页数，通过计算得到
	TotalRecord int64      // 总记录数，通过查询数据库得到
}

// IsHasPrev 判断是否存在上一页
func (p *PageProduct) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 判断是否存在下一页
func (p *PageProduct) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取上一页的页码
func (p *PageProduct) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

// GetNextPageNo 获取下一页的页码
func (p *PageProduct) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}

// GetPageProducts 获取产品列表分页结构
func GetPageProducts(pageNo string) (*PageProduct, error) {
	query := "select count(*) from products"
	query2 := "select id, category_id, name, price, stock, sales, img_path, detail, hot_point from products"
	query2 += " limit ?, ?"

	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64) // 当前页页码，转换为int类型
	var pageSize int64 = 12                        // 每页记录数
	var totalRecord int64                          // 总记录数
	var totalPageNo int64                          // 总页数

	// 查询数据库获取总记录数
	err := util.Db.QueryRow(query).Scan(&totalRecord)
	if err != nil {
		return nil, err
	}

	// 计算总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	// 获取当前页产品列表，返回分页结构
	rows, err := util.Db.Query(query2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	// 产品列表
	var ps []*Product
	for rows.Next() {
		p := &Product{}
		rows.Scan(&p.ID, &p.Category_id, &p.Name, &p.Price, &p.Stock, &p.Sales,
			&p.ImgPath, &p.Detail, &p.HotPoint)
		ps = append(ps, p)
	}

	// 分页结构
	pageProduct := &PageProduct{
		Products:    ps,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return pageProduct, nil
}
