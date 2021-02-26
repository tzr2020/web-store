package model

import (
	"fmt"
	"log"
	"testing"
)

func TestProduct(t *testing.T) {
	// t.Run("获取所有产品", testGetProducts)
	// t.Run("获取产品，根据产品id", testGetProduct)
	// t.Run("修改产品的库存和销量，根据产品id", testUpdateStockAndSales)
}

func testGetProducts(t *testing.T) {
	fmt.Println("测试获取所有产品")

	ps, err := GetProducts()
	if err != nil {
		log.Println(err)
		return
	}

	for k, v := range ps {
		fmt.Println(k+1, v)
	}
}

func testGetProduct(t *testing.T) {
	fmt.Println("测试获取产品，根据产品id")

	pid := "1"
	p, err := GetProduct(pid)
	if err != nil {
		log.Printf("数据库错误: %v\n", err)
		return
	}

	fmt.Printf("产品信息: %v\n", p)
}

func testUpdateStockAndSales(t *testing.T) {
	fmt.Println("测试修改产品的库存和销量，根据产品id")

	product := &Product{
		ID:    2,
		Stock: 100,
		Sales: 0,
	}

	if err := product.UpdateStockAndSales(); err != nil {
		log.Printf("数据库更新产品的库存和销量发生错误: %v\n", err)
		return
	}
}
