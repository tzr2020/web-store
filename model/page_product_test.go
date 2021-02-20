package model

import (
	"fmt"
	"log"
	"testing"
)

func TestPageProduct(t *testing.T) {
	// t.Run("获取产品列表，根据分页", testGetPageProducts)
	// t.Run("测试获取产品列表，根据分页和价格区间", testGetPageProductsByPrice)
}

func testGetPageProducts(t *testing.T) {
	fmt.Println("测试获取产品列表，根据分页")

	PageProduct, err := GetPageProducts("1")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("当前页码：", PageProduct.PageNo)
	fmt.Println("总页数：", PageProduct.TotalPageNo)
	fmt.Println("总记录数：", PageProduct.TotalRecord)
	fmt.Println("当前页产品列表：")

	for _, v := range PageProduct.Products {
		fmt.Println(v)
	}
}

func testGetPageProductsByPrice(t *testing.T) {
	fmt.Println("测试获取产品列表，根据分页和价格区间")

	PageProduct, err := GetPageProductsByPrice("1", "50", "80")
	if err != nil {
		log.Println("查询数据库错误：", err)
		return
	}

	fmt.Println("当前页码：", PageProduct.PageNo)
	fmt.Println("总页数：", PageProduct.TotalPageNo)
	fmt.Println("总记录数：", PageProduct.TotalRecord)
	fmt.Println("当前页产品列表：")

	for _, v := range PageProduct.Products {
		fmt.Println(v)
	}
}
