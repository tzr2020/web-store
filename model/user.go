package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"web-store/util"
)

var (
	// ErrUserNotFound used when the user wasn't found on the db.
	// 在数据库没有找到用户。
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

// CheckUsernameAndPassword 从数据库验证用户名和密码，返回用户
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

// CheckUsername 从数据库验证用户名，返回用户
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

// AddUser 数据库新增用户
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

// GetUsersPage 后台查询数据库获取当前页会员用户
func GetUserPage(pageNo string, pageSize string) ([]*User, error) {
	query := "select id, username, password, email, nickname, sex, avatar, phone, country, province, city from users limit ?, ?"

	// 数据类型转换
	intPageNo, err := strconv.Atoi(pageNo) // 当前页页码
	if err != nil {
		return nil, err
	}
	intPageSize, err := strconv.Atoi(pageSize) // 每页记录数
	if err != nil {
		return nil, err
	}

	// 查询数据库获取当前页数据
	rows, err := util.Db.Query(query, (intPageNo-1)*intPageSize, intPageSize)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Nickname,
			&user.Sex, &user.Avatar, &user.Phone, &user.Country, &user.Province, &user.City)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Add 后台从数据库添加会员用户
func (user *User) Add() error {
	// SQL语句
	query := `insert into users(username, password, email, nickname, sex, phone, 
		country, province, city) values(?,?,?,?,?,?,?,?,?)`

	// 预编译SQL语句
	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return err
	}

	// 执行SQL语句
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Nickname,
		user.Sex, user.Phone, user.Country, user.Province, user.City)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser 后台从数据库删除会员用户，根据用户id
func DeleteUser(id int) error {
	query := `delete from users where id = ?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// Update 后台从数据库更新会员用户
func (user User) Update() error {
	query := `update users set username=?, password=?, email=?, nickname=?, sex=?, 
		phone=?, country=?, province=?, city=? where id=?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&user.Username, &user.Password, &user.Email, &user.Nickname,
		&user.Sex, &user.Phone, &user.Country, &user.Province, &user.City, &user.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserAvatar 数据库更新用户头像
func UpdateUserAvatar(avatar string, uid string) (err error) {
	query := `update users set avatar = ? where id = ?`

	stmt, err := util.Db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(avatar, uid)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(pageNo int, pageSize int, uid int) ([]*User, error) {
	query := `select id, username, password, email, nickname, sex, avatar, phone, country, province, city
		from users
		where id=?
		limit ?,?`

	// 查询数据库获取当前页数据
	rows, err := util.Db.Query(query, uid, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Nickname,
			&user.Sex, &user.Avatar, &user.Phone, &user.Country, &user.Province, &user.City)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
