package model

import (
	"fmt"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	// t.Run("验证用户名和密码", testCheckUsernameAndPassword)
	// t.Run("验证用户名", testCheckUsername)
	// t.Run("新增用户", testAddUser)
	// t.Run("后台获取所有会员用户", testGetUserPage)
	// t.Run("后台添加会员用户", testAddUser2)
	t.Run("后台删除会员用户，根据用户id", testDeleteUser)
}

func testCheckUsernameAndPassword(t *testing.T) {
	fmt.Println("验证用户名和密码")

	username := "user"
	password := "12345678"

	user, err := CheckUsernameAndPassword(username, password)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("用户信息：", user)
}

func testCheckUsername(t *testing.T) {
	fmt.Println("验证用户名")

	username := "user"

	user, err := CheckUsername(username)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("用户信息：", user)
}

func testAddUser(t *testing.T) {
	fmt.Println("新增用户")

	user := &User{
		Username: "user3",
		Password: "12345678",
		Email:    "user3@qq.com",
	}

	err := AddUser(user)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("新增用户成功")
}

func testGetUserPage(t *testing.T) {
	fmt.Println("后台获取所有会员用户")

	users, err := GetUserPage("1", "10")
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range users {
		fmt.Println(v)
	}
}

func testAddUser2(t *testing.T) {
	user := &User{
		Username: "test",
		Password: "123456",
		Email:    "test@qq.com",
		Nickname: "test",
		Sex:      0,
		Phone:    "13699567845",
		Country:  "中国",
		Province: "广东",
		City:     "深圳",
	}

	err := user.Add()
	if err != nil {
		log.Println("后台添加会员用户失败")
		return
	}

	fmt.Println("后台添加会员用户成功")
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser(38)
	if err != nil {
		log.Println("后台删除会员用户失败")
		return
	}

	log.Println("后台删除会员用户成功")
}
