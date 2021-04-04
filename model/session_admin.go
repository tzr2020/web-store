package model

import (
	"web-store/util"
)

type AdminSession struct {
	SessionID string `json:"session_id"`
	AdminID   int    `json:"admin_id,string"`
	Admin     *Admin `json:"admin"`
}

func (as AdminSession) Add() error {
	query := `insert into session_admin values(?, ?)`

	_, err := util.Db.Exec(query, as.SessionID, as.AdminID)
	if err != nil {
		return err
	}

	return nil
}

// GetAdminSession 查询数据库，获取管理员session，根据sessionID
func GetAdminSession(sessID string) (as AdminSession, err error) {
	query := `select session_id, admin_id from session_admin where session_id=?`

	err = util.Db.QueryRow(query, sessID).Scan(&as.SessionID, &as.AdminID)
	if err != nil {
		return
	}

	return
}

// GetAdminBySessID 查询数据库，获取管理员，根据sessionID
func GetAdminBySessID(sessID string) (a Admin, err error) {
	query := `select admin_id from session_admin where session_id=?`
	var adminID int
	err = util.Db.QueryRow(query, sessID).Scan(&adminID)
	if err != nil {
		return
	}
	a, err = GetAdminByID(adminID)
	if err != nil {
		return
	}
	return
}

// DeleteAdminSessionByID 数据库生成AdminSession，根据SessionID
func DeleteAdminSessionByID(sessID string) (err error) {
	query := `delete from session_admin where session_id=?`
	_, err = util.Db.Exec(query, sessID)
	if err != nil {
		return
	}
	return
}

// GetAdminSessionPage 查询数据库，获取管理员Session列表，根据当前页的页码和每页记录条数
func GetAdminSessionPage(pageNo int, pageSize int) ([]*AdminSession, error) {
	query := `select session_id, admin_id
		from session_admin 
		limit ?, ?`

	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var ss []*AdminSession
	for rows.Next() {
		s := &AdminSession{}
		err = rows.Scan(&s.SessionID, &s.AdminID)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}

	return ss, nil
}

func (s AdminSession) Delete() error {
	query := `delete from session_admin 
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
