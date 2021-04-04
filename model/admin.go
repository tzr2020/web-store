package model

import (
	"web-store/util"
)

type Admin struct {
	ID       int    `json:"id,string"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckUsernameAndPassword 查询数据库，获取管理员id，根据登录名称和密码
func (a Admin) CheckUsernameAndPassword() (admin Admin, err error) {
	query := `select id from admins where username=? and password=?`

	err = util.Db.QueryRow(query, a.Username, a.Password).Scan(&admin.ID)
	if err != nil {
		return
	}

	return
}

// GetAdminByID 查询数据库，获取管理员，根据adminID
func GetAdminByID(adminID int) (a Admin, err error) {
	query := `select id, username, password from admins where id=?`
	err = util.Db.QueryRow(query, adminID).Scan(&a.ID, &a.Username, &a.Password)
	if err != nil {
		return
	}
	return
}

// GetAdminPage 后台查询数据库获取当前页管理员
func GetAdminPage(pageNo int, pageSize int) ([]*Admin, error) {
	query := `select id, username, password
		from admins
		limit ?, ?`

	// 查询数据库获取当前页数据
	rows, err := util.Db.Query(query, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var as []*Admin
	for rows.Next() {
		a := &Admin{}
		err = rows.Scan(&a.ID, &a.Username, &a.Password)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}

	return as, nil
}

func (a Admin) Add() error {
	// SQL语句
	query := `insert into admins (username, password)
		values(?,?)`

	// 预编译SQL语句
	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	// 执行SQL语句
	_, err = stmt.Exec(&a.Username, &a.Password)
	if err != nil {
		return err
	}

	return nil
}

func (a Admin) Delete() error {
	query := `delete from admins
		where id = ?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&a.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a Admin) Update() error {
	query := `update admins
		set username=?, password=?
		where id=?`

	stmt, err := util.Db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&a.Username, &a.Password, &a.ID)
	if err != nil {
		return err
	}

	return nil
}
