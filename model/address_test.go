package model

import (
	"fmt"
	"log"
	"testing"
)

func TestAddress(t *testing.T) {
	t.Run("测试获取会员用户的所有收货地址，根据用户id", testGetAddress)
}

func testGetAddress(t *testing.T) {
	fmt.Println("测试获取会员用户的所有收货地址，根据用户id")

	addresses, err := GetAddressByUserID(1)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range addresses {
		fmt.Println("收货地址:", v)
	}
}
