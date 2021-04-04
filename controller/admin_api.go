package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web-store/model"
	"web-store/util"
)

func registerAPIAdminRouters() {
	http.HandleFunc("/api/admin/login", login)
	http.HandleFunc("/api/admin/username", getAdminUsername)
	http.HandleFunc("/api/admin/logout", logout)
	http.HandleFunc("/api/admin/list", adminList)
	http.HandleFunc("/api/admin", admin)
}

// AdminIsLoginByCookie 通过cookie来判断管理员是否已经登录
func AdminIsLoginByCookie(r *http.Request) (ok bool, as model.AdminSession, err error) {
	// 获取请求中指定名称的cookie
	cookie, err := r.Cookie("admin")
	if err != nil {
		return
	}
	// 根据cookie查询数据库，判断管理员是否已经登录
	sessID := cookie.Value
	as, err = model.GetAdminSession(sessID)
	if err != nil {
		return
	}
	if as.AdminID < 1 {
		return
	}
	return true, as, nil
}

// login 后台管理员登录处理器
func login(w http.ResponseWriter, r *http.Request) {

	// 通过cookie判断管理员是否已经登录
	ok, _, err := AdminIsLoginByCookie(r)
	if err != nil && err != http.ErrNoCookie && err != sql.ErrNoRows {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 管理员已经登录，页面重定向到后台管理系统首页
	if ok {
		w.Header().Set("Location", "/manage/index")
		w.WriteHeader(302)
		return
	}

	// 根据不同的请求方法，对请求进行不同的处理
	switch r.Method {
	case http.MethodGet:
		// 返回管理员登录页面
		adminLoginPage(w)

	case http.MethodPost:
		// 解码请求包体中的JSON数据到结构体
		dec := json.NewDecoder(r.Body)
		admin := model.Admin{}
		err := dec.Decode(&admin)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 查询数据库，验证管理员的登录名称和密码
		admin, err = admin.CheckUsernameAndPassword()
		if err == sql.ErrNoRows {
			util.ResponseWriteJson(w, util.Json{
				Code: 401,
				Msg:  "管理员登录失败，登录名称或密码错误",
			})
			return
		}
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		if admin.ID < 1 {
			util.ResponseWriteJson(w, util.Json{
				Code: 401,
				Msg:  "管理员登录失败，登录名称或密码错误",
			})
			return
		}
		// 创建session，并添加到数据库
		sess := model.AdminSession{
			SessionID: util.CreateUUID(),
			AdminID:   admin.ID,
		}
		err = sess.Add()
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		// 响应设置cookie
		http.SetCookie(w, &http.Cookie{
			Name:   "admin",
			Value:  sess.SessionID,
			MaxAge: 60 * 60 * 24 * 7,
			Path:   "/",
		})
		// 返回JSON响应
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "管理员登录成功",
		})
	}
}

// adminLoginPage 执行模板，返回管理员登录页面
func adminLoginPage(w http.ResponseWriter) {
	util.ExecuteTpl(w, "./view/template/manage/admin-login.html")
}

// getAdminUsername 返回包含管理员用户名数据的JSON响应
func getAdminUsername(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	c, err := r.Cookie("admin")
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 查询数据库，获取管理员
	sessID := c.Value
	admin, err := model.GetAdminBySessID(sessID)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 返回成功的JSON响应
	util.ResponseWriteJson(w, util.Json{
		Code: 200,
		Msg:  "获取管理员用户名成功",
		Data: admin.Username,
	})
}

// 管理员退出账号处理器
func logout(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	c, err := r.Cookie("admin")
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 数据库删除AdminSession
	sessID := c.Value
	err = model.DeleteAdminSessionByID(sessID)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}
	// 将cookie设置为失效并发送
	c.Path = "/"
	c.MaxAge = -1
	http.SetCookie(w, c)
	// 重定向到管理员登录页面
	w.Header().Set("Location", "/api/admin/login")
	w.WriteHeader(302)
}

// adminList 获取管理员列表，根据当前页的页码和每页记录条数
func adminList(w http.ResponseWriter, r *http.Request) {
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
		count, err := util.GetJsonDataCount("admins")
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 查询数据库，获取列表，根据当前页的页码和每页记录条数
		list, err := model.GetAdminPage(intPageNo, intPageSize)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		// 将JSON结构编码为JSON格式数据后写入响应
		j := util.Json{
			Code:  200,
			Msg:   "获取管理员列表成功",
			Count: count,
			Data:  list,
		}
		util.ResponseWriteJson(w, j)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	// 解码包体JSON数据到结构
	dec := json.NewDecoder(r.Body)
	a := model.Admin{}
	err := dec.Decode(&a)
	if err != nil {
		log.Println(err)
		util.ResponseWriteJsonOfInsideServer(w)
		return
	}

	switch r.Method {
	case http.MethodPost:
		// 添加管理员

		// 数据库添加管理员
		if err = a.Add(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库添加管理员成功",
		})
	case http.MethodPut:
		// 更新管理员

		// 数据库更新管理员
		if err = a.Update(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库更新管理员成功",
		})
	case http.MethodDelete:
		// 删除管理员

		// 数据库删除管理员
		if err = a.Delete(); err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}
		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "数据库删除管理员成功",
		})
	}
}
