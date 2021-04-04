package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIProductRouters() {
	http.HandleFunc("/api/products", products)
	http.HandleFunc("/api/product", product)
	http.HandleFunc("/api/uploadProductImg", uploadProductImg)
	http.HandleFunc("/api/uploadProductDetail", uploadProductDetail)
}

// products 获取产品列表，根据当前页的页码和每页记录条数
func products(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("products")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetProductPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取产品列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func product(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	p := model.Product{}
	err := dec.Decode(&p)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加产品

		// 数据库添加产品
		if err = p.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加产品成功",
		})
	case http.MethodPut:
		// 更新产品

		// 数据库更新产品
		if err = p.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新产品成功",
		})
	case http.MethodDelete:
		// 删除产品

		// 数据库删除产品
		if err = p.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除产品成功",
		})
	}
}

// uploadProduct 上传产品图片
func uploadProductImg(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		productID := r.PostFormValue("productID")

		// 解析请求包体，设置用于存储包体数据的内存大小
		r.ParseMultipartForm(32 << 20)

		// 获取上传文件
		file, handler, err := r.FormFile("productImg")
		defer file.Close()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 获取上传文件后缀名
		var suffix string
		reg, err := regexp.Compile(`.(jpg|png)$`)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		matches := reg.FindStringSubmatch(handler.Filename)
		if len(matches) > 0 {
			suffix = matches[0]
		} else {
			util.ResponseWriteJson(w, util.Json{
				Code: 500,
				Msg:  "上传产品图片失败，图片文件后缀名要求是.jpg或.png",
			})
			return
		}

		// 存储上传文件本地
		path := "view/static/img/product-img/" + productID + suffix
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 更新数据库
		imgPath := "static/img/product-img/" + productID + suffix
		err = model.UpdateProductImg(imgPath, productID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "上传产品图片成功",
		})
	}
}

// uploadProductDetail 上传产品详情
func uploadProductDetail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		productID := r.PostFormValue("productID")

		// 解析请求包体，设置用于存储包体数据的内存大小
		r.ParseMultipartForm(32 << 20)

		// 获取上传文件
		file, handler, err := r.FormFile("productDetail")
		defer file.Close()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 获取上传文件后缀名
		var suffix string
		reg, err := regexp.Compile(`.(jpg|png)$`)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		matches := reg.FindStringSubmatch(handler.Filename)
		if len(matches) > 0 {
			suffix = matches[0]
		} else {
			util.ResponseWriteJson(w, util.Json{
				Code: 500,
				Msg:  "上传产品图片失败，图片文件后缀名要求是.jpg或.png",
			})
			return
		}

		// 存储上传文件本地
		path := "view/static/img/product-detail/" + productID + suffix
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 更新数据库
		detail := "static/img/product-detail/" + productID + suffix
		err = model.UpdateProductDetail(detail, productID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "上传产品详情成功",
		})
	}
}
