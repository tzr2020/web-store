package model

import (
	"fmt"
	"log"
	"testing"
)

func TestCategory(t *testing.T) {
	// t.Run("获取产品类别", testGetCategory)
	// t.Run("获取所有产品类别", testGetCategories)
}

func testGetCategory(t *testing.T) {
	fmt.Println("测试获取产品类别")

	cate_id := "1"
	cate, err := GetCategory(cate_id)

	if err != nil {
		log.Printf("数据库错误: %v\n", err)
		return
	}

	fmt.Printf("产品类别信息: %v\n", cate)
}

func testGetCategories(t *testing.T) {
	fmt.Println("测试获取所有产品类别")

	cates, err := GetCategories()

	if err != nil {
		log.Printf("数据库错误: %v\n", err)
		return
	}

	for _, v := range cates {
		fmt.Printf("产品类别信息: %v\n", v)
	}
}
