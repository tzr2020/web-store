package model

import (
	"fmt"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	// t.Run("验证用户名和密码", testCheckUsernameAndPassword)
	// t.Run("验证用户名", testCheckUsername)
	t.Run("新增用户", testAddUser)
}

func testCheckUsernameAndPassword(t *testing.T) {
	fmt.Println("验证用户名和密码")

	username := "user"
	password := "12345678"

	user, err := CheckUsernameAndPassword(username, password)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("用户信息：", user)
}

func testCheckUsername(t *testing.T) {
	fmt.Println("验证用户名")

	username := "user"

	user, err := CheckUsername(username)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("用户信息：", user)
}

func testAddUser(t *testing.T) {
	fmt.Println("新增用户")

	user := &User{
		Username: "user4",
		Password: "12345678",
		Email:    "user4@qq.com",
		Phone:    "11122233344",
		State:    1,
	}

	err := AddUser(user)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("新增用户成功")
}
