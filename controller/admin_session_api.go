package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIAdminSessionRouters() {
	http.HandleFunc("/api/admin/session/list", adminSessionList)
	http.HandleFunc("/api/admin/session", adminSession)
}

func adminSessionList(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串获取数据
	pageNo := r.FormValue("pageNo")     // 当前页页码
	pageSize := r.FormValue("pageSize") // 每页记录条数
	// 数据类型转换
	intPageNo, err := strconv.Atoi(pageNo)
	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// 查询数据库，获取列表的总记录条数
		count, err := util.GetJsonDataCount("session_admin")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetAdminSessionPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取管理员Session列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func adminSession(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	s := model.AdminSession{}
	err := dec.Decode(&s)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		// 删除管理员Session

		// 数据库删除管理员Session
		if err = s.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除管理员Session成功",
		})
	}
}
