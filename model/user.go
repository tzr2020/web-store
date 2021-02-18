package model

import (
	"database/sql"
	"errors"
	"fmt"
	"web-store/util"
)

var (
	// ErrUserNotFound used when the user wasn't found on the db.
	// 在数据库没有找到用户。
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Nickname string
	sex      string
	avatar   string
	phone    string
	country  string
	province string
	city     string
}

// CheckUsernameAndPassword 验证用户名和密码
func CheckUsernameAndPassword(username string, password string) (*User, error) {
	user := &User{}
	query := "select id, username, password, email from users where username=? and password=?"

	// 执行SQL，得到查询结果
	row := util.Db.QueryRow(query, username, password)
	// 将查询结果扫描到结构体
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err == sql.ErrNoRows {
		return user, ErrUserNotFound
	}

	if err != nil {
		return user, fmt.Errorf("could not query select user: %v", err)
	}

	return user, nil
}

// CheckUsername 验证用户名
func CheckUsername(username string) (user *User, err error) {
	user = &User{}
	query := "select id, username, password, email from users where username=?"

	err = util.Db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Email)

	if err == sql.ErrNoRows {
		return user, ErrUserNotFound
	}

	return
}

// AddUser 新增用户
func AddUser(user *User) error {
	query := "insert into users (username, password, email) values(?,?,?)"

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}
