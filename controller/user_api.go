package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIUserRoutes() {
	http.HandleFunc("/api/users", Users)
	http.HandleFunc("/api/user", User)
}

func Users(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串获取数据
	pageNo := r.FormValue("pageNo")
	pageSize := r.FormValue("pageSize")

	switch r.Method {
	case http.MethodGet:
		// 获取用户列表，根据当前页的页码和记录数

		// 查询数据库，获取用户列表
		users, err := model.GetUserPage(pageNo, pageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，获取用户表的记录条数
		count, err := util.GetJsonDataCount("users")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 准备JSON结构
		j := util.Json{
			Code:  200,
			Msg:   "成功",
			Count: count,
			Data:  users,
		}
		// 将JSON结构编码为JSON数据格式后写入响应
		enc := json.NewEncoder(w)
		err = enc.Encode(j)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJson(w, j)
		}
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// 添加用户

		// 从HTML表单获取数据
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		nickname := r.PostFormValue("nickname")
		sex := r.PostFormValue("sex")
		phone := r.PostFormValue("phone")
		country := r.PostFormValue("country")
		province := r.PostFormValue("province")
		city := r.PostFormValue("city")
		// 数据类型转换
		intSex, err := strconv.Atoi(sex)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 准备JSON结构
		j := util.Json{}

		// 向数据库添加用户
		user := &model.User{
			Username: username,
			Password: password,
			Email:    email,
			Nickname: nickname,
			Sex:      intSex,
			Phone:    phone,
			Country:  country,
			Province: province,
			City:     city,
		}
		err = user.Add()
		if err != nil {
			log.Println(err)
			j.Code = http.StatusInternalServerError
			j.Msg = "数据库添加会员用户失败"
			util.ResponseWriteJson(w, j)
			return
		}
		j.Code = http.StatusOK
		j.Msg = "数据库添加会员用户成功"
		util.ResponseWriteJson(w, j)

	case http.MethodDelete:
		// 删除用户

		// 解码包体JSON数据到结构
		dec := json.NewDecoder(r.Body)
		user := model.User{}
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 准备JSON结构
		j := util.Json{}

		// 数据库删除会员用户
		err = model.DeleteUser(user.ID)
		if err != nil {
			log.Println(err)
			j.Code = 500
			j.Msg = "数据库删除会员用户失败"
			util.ResponseWriteJson(w, j)
			return
		}
		j.Code = 200
		j.Msg = "数据库删除会员用户成功"
		util.ResponseWriteJson(w, j)

	case http.MethodPut:
		// 更新用户

		// 从HTML表单获取数据
		id := r.PostFormValue("id")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		nickname := r.PostFormValue("nickname")
		sex := r.PostFormValue("sex")
		phone := r.PostFormValue("phone")
		country := r.PostFormValue("country")
		province := r.PostFormValue("province")
		city := r.PostFormValue("city")
		// 数据类型转换
		intID, err := strconv.Atoi(id)
		intSex, err := strconv.Atoi(sex)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 准备JSON结构
		j := util.Json{}

		// 向数据库更新用户
		user := &model.User{
			ID:       intID,
			Username: username,
			Password: password,
			Email:    email,
			Nickname: nickname,
			Sex:      intSex,
			Phone:    phone,
			Country:  country,
			Province: province,
			City:     city,
		}
		err = user.Update()
		if err != nil {
			log.Println(err)
			j.Code = http.StatusInternalServerError
			j.Msg = "数据库更新会员用户失败"
			util.ResponseWriteJson(w, j)
			return
		}
		j.Code = http.StatusOK
		j.Msg = "数据库更新会员用户成功"
		util.ResponseWriteJson(w, j)
	}
}
