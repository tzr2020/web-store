package model

import (
	"fmt"
	"log"
	"testing"
)

func TestProduct(t *testing.T) {
	// t.Run("获取所有产品", testGetProducts)
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
