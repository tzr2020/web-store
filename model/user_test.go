package model

import (
	"fmt"
	"log"
	"testing"
)

func TestCheckUsernameAndPassword(t *testing.T) {
	fmt.Println("验证用户名和密码")

	username := "user"
	password := "12345678"

	user, err := CheckUsernameAndPassword(username, password)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("用户信息：", user)
}

func TestCheckUsername(t *testing.T) {
	fmt.Println("验证用户名")

	username := "user"

	user, err := CheckUsername(username)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("用户信息：", user)
}
