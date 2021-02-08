package model

import (
	"web-store/util"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Phone    string
	State    int
}

// CheckUsernameAndPassword 验证用户名和密码
func CheckUsernameAndPassword(username string, password string) (user *User, err error) {
	user = &User{}
	sql := "select id, username, password, email, phone, state from user where username=? and password=?"

	// 执行SQL，得到查询结果
	row := util.Db.QueryRow(sql, username, password)
	// 将查询结果扫描到结构体
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.State)

	return
}

// CheckUsername 验证用户名
func CheckUsername(username string) (user *User, err error) {
	user = &User{}
	sql := "select id, username, password, email, phone, state from user where username=?"

	err = util.Db.QueryRow(sql, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.State)

	return
}

// CheckUsername 验证用户账号状态
func CheckUserState() {
	// todo
}

// AddUser 新增用户
func AddUser(user *User) error {
	sql := "insert into user(username, password, email, phone, state) values(?,?,?,?,?)"

	stmt, err := util.Db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Phone, user.State)
	if err != nil {
		return err
	}

	return nil
}
