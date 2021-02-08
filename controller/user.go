package controller

import (
	"html/template"
	"log"
	"net/http"
	"web-store/model"
)

// 注册路由
func registerUserRoutes() {
	http.HandleFunc("/user/login", handlerLogin)
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	// GET请求
	if r.Method == http.MethodGet {
		// 返回登录页面
		getLoginPage(w, "")
	}

	// POST请求
	if r.Method == http.MethodPost {
		// 从表单获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		// 从数据库获取用户信息
		user, err := model.CheckUsernameAndPassword(username, password)
		if err != nil {
			log.Println(err)
		}

		// 验证登录
		CheckLogin(user, w)
	}
}

// 返回登录页面
func getLoginPage(w http.ResponseWriter, msg string) {
	t, err := template.ParseFiles("./view/page/user/login.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, msg)
	if err != nil {
		log.Println(err)
	}
}

// CheckLogin 验证登录
func CheckLogin(user *model.User, w http.ResponseWriter) {
	// 登录成功，返回欢迎用户页面
	if user.ID > 0 {
		t, err := template.ParseFiles("./view/page/user/login_success.html")
		if err != nil {
			log.Println(err)
		}

		err = t.Execute(w, user.Username)
		if err != nil {
			log.Println(err)
		}
	} else {
		// 登录失败，返回登录页面
		getLoginPage(w, "用户名或密码错误")
	}
}
