package model

import (
	"fmt"
	"testing"
)

func TestCartItem(t *testing.T) {
	// t.Run("测试获取购物项，根据购物车id", testGetCartItemByCartID)
	// t.Run("测试获取购物项，根据购物车id和产品id", testGetCartItemByCartIDAndProductID)
}

func testGetCartItemByCartID(t *testing.T) {
	fmt.Println("测试获取购物项，根据购物车id")

	cItems, err := GetCartItemByCartID("5103ce2c-9940-44f1-5ca6-53c2997c39bb")
	if err != nil {
		return
	}

	for _, v := range cItems {
		fmt.Printf("购物项信息: %v\n", v)
	}
}

func testGetCartItemByCartIDAndProductID(t *testing.T) {
	fmt.Println("测试获取购物项，根据购物车id和产品id")

	cItem, err := GetCartItemByCartIDAndProductID("5103ce2c-9940-44f1-5ca6-53c2997c39bb", "1")
	if err != nil {
		return
	}

	fmt.Printf("购物项信息: %v\n", cItem)
}
