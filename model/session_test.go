package model

import (
	"fmt"
	"log"
	"testing"
	"web-store/util"
)

func TestSession(t *testing.T) {
	// t.Run("新增Session", testAddSession)
	// t.Run("删除Session，根据SessionID", testDeteleSession)
	// t.Run("获取Session，根据SessionID", testGetSession)
}

func testAddSession(t *testing.T) {
	fmt.Println("新增Session")

	uuid := util.CreateUUID()
	sess := &Session{
		SessionID: uuid,
		Username:  "user",
		UserID:    1,
	}

	if err := AddSession(sess); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("新增Session成功")
}

func testDeteleSession(t *testing.T) {
	fmt.Println("删除Session，根据SessionID")

	err := DeleteSession("e3e38aef-ad04-4e2a-6727-e1e5dadea15a")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("删除Session成功")
}

func testGetSession(t *testing.T) {
	fmt.Println("获取Session，根据SessionID")

	sess, err := GetSession("6bae8cb0-1508-4783-4491-d8b1964ea181")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Session：", sess)
}
