package model

import (
	"database/sql"
	"fmt"
	"web-store/util"
)

type Session struct {
	SessionID         string
	Username          string
	UserID            int
	PageProduct       *PageProduct        // 用于模板
	Product           *Product            // 用于模板
	Cart              *Cart               // 用于模板
	Order             *Order              // 用于模板
	OrderPaymentTypes []*OrderPaymentType // 用于模板
	Address           *Address            // 用于模板
}

// AddSession 数据库新增Session
func AddSession(sess *Session) error {
	query := "insert into session values (?, ?, ?)"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return fmt.Errorf("准备SQL语句错误：%v", err)
	}

	_, err = stmt.Exec(sess.SessionID, sess.Username, sess.UserID)
	if err != nil {
		return fmt.Errorf("执行SQL语句错误：%v", err)
	}

	return nil
}

// DeleteSession 根据SessionID，数据库删除Session
func DeleteSession(sessID string) error {
	query := "delete from session where session_id = ?"

	sess, err := GetSession(sessID)
	if err != nil {
		return err
	}

	_, err = util.Db.Exec(query, sess.SessionID)
	if err != nil {
		return fmt.Errorf("执行SQL语句错误：%v", err)
	}

	return nil
}

// GetSession 根据SessionID，从数据库获取Session
func GetSession(sessID string) (*Session, error) {
	query := "select session_id, username, user_id from session where session_id = ?"

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, fmt.Errorf("准备SQL语句错误：%v", err)
	}

	sess := &Session{}

	err = stmt.QueryRow(sessID).Scan(&sess.SessionID, &sess.Username, &sess.UserID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("执行SQL语句错误：%v", err)
	}

	return sess, nil
}
