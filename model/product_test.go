package model

import (
	"fmt"
	"log"
	"testing"
)

func TestProduct(t *testing.T) {
	// t.Run("获取所有产品", testGetProducts)
	// t.Run("获取产品，根据产品id", testGetProduct)
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
