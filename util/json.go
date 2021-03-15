package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// Json 用于返回响应的JSON结构
type Json struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

// GetJsonDataCount 获取数据库某张表的记录条数
func GetJsonDataCount(databaseTableName string) (count int, err error) {
	query := "select count(*) from " + databaseTableName

	err = Db.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}

	return
}

// ResponseWriteJson 将JSON结构编码为JSON格式的数据后写入响应
func ResponseWriteJson(w http.ResponseWriter, j Json) {
	enc := json.NewEncoder(w)
	err = enc.Encode(j)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrServerInside.Error()))
	}
}

// ResponseWriteJsonOfInsideServer 将表示服务器内部发生错误的JSON结构编码为JSON格式的数据后写入响应
func ResponseWriteJsonOfInsideServer(w http.ResponseWriter) {
	j := Json{
		Code: http.StatusInternalServerError,
		Msg:  ErrServerInside.Error(),
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(j)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrServerInside.Error()))
	}
}
