package model

import (
	"database/sql"
	"fmt"
	"web-store/util"
)

type Session struct {
	SessionID         string              `json:"session_id"`
	Username          string              `json:"username"`
	UserID            int                 `json:"user_id,string"`
	PageProduct       *PageProduct        `json:"page_product"`        // 用于模板
	Product           *Product            `json:"product"`             // 用于模板
	Cart              *Cart               `json:"cart"`                // 用于模板
	Order             *Order              `json:"order"`               // 用于模板
	Orders            []*Order            `json:"orders"`              // 用于模板
	OrderPaymentTypes []*OrderPaymentType `json:"order_payment_types"` // 用于模板
	Address           *Address            `json:"address"`             // 用于模板
	OrderItems        []*OrderItem        `json:"order_items"`         // 用于模板
	OrderAddress      *OrderAddress       `json:"order_address"`       // 用于模板
	IndexPage         *IndexPage          `json:"index_page"`          // 用于模板
	Nav               *Nav                `json:"nav"`                 // 用于模板
	Category          *Category           `json:"category"`            // 用于模板
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

// GetUserSessionPage 查询数据库，获取用户Session列表，根据当前页的页码和每页记录条数
func GetUserSessionPage(pageNo int, pageSize int) ([]*Session, error) {
	query := `select session_id, username, user_id
		from session 
		limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var ss []*Session
	for rows.Next() {
		s := &Session{}
		err = rows.Scan(&s.SessionID, &s.Username, &s.UserID)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}

	return ss, nil
}

func (s Session) Delete() error {
	query := `delete from session 
		where session_id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&s.SessionID)
	if err != nil {
		return err
	}

	return nil
}
