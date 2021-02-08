package controller

import (
	"html/template"
	"log"
	"net/http"
	"web-store/model"
)

// registerUserRoutes 路由器注册用户相关的路由规则和处理器
func registerUserRoutes() {
	http.HandleFunc("/user/login", handlerLogin)
	http.HandleFunc("/user/regist", handlerRegist)
	http.HandleFunc("/user/checkUsername", handlerCheckUsername)
}

// handlerLogin 用户登录处理器
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
			log.Println("CheckUsernameAndPassword: ", err)
		}

		// 验证登录
		CheckLogin(user, w)
	}
}

// getLoginPage 模板引擎生成最终页面，并返回登录页面
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
	if user.ID > 0 {
		// 登录成功，模板引擎生成最终页面，并返回欢迎用户页面
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

func handlerRegist(w http.ResponseWriter, r *http.Request) {
	// GET请求
	if r.Method == http.MethodGet {
		// 返回注册页面
		getRegistPage(w, "")
	}

	// POST请求
	if r.Method == http.MethodPost {
		// 从表单获取用户注册信息
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		phone := r.PostFormValue("phone")

		newUser := &model.User{
			Username: username,
			Password: password,
			Email:    email,
			Phone:    phone,
			State:    1,
		}

		// 验证注册
		CheckRegist(w, newUser)
	}

}

// getRegistPage 模板引擎生成最终页面，并返回注册页面
func getRegistPage(w http.ResponseWriter, msg string) {
	t, err := template.ParseFiles("./view/page/user/regist.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, msg)
	if err != nil {
		log.Println(err)
	}
}

// 验证注册信息
func CheckRegist(w http.ResponseWriter, newUser *model.User) {
	// 从数据库查询是否存在某用户名的用户，判断用户名是否已存在
	user, err := model.CheckUsername(newUser.Username)
	if err != nil {
		log.Println("CheckUsername: ", err)
	}

	// 用户名已存在，不可用
	if user.ID > 0 {
		// 返回注册页面，并带上提示信息
		getRegistPage(w, newUser.Username+"用户名已存在！")
	} else {
		// 用户名可用

		// 数据库新增用户
		err = model.AddUser(newUser)
		if err != nil {
			log.Println(err)
		}

		// 返回注册成功页面
		t, err := template.ParseFiles("./view/page/user/regist_success.html")
		if err != nil {
			log.Println(err)
		}

		err = t.Execute(w, newUser.Username)
		if err != nil {
			log.Println(err)
		}
	}
}

// handlerCheckUsername 验证用户名处理器
func handlerCheckUsername(w http.ResponseWriter, r *http.Request) {
	// 从Ajax请求里获取用户输入的用户名
	username := r.PostFormValue("username")

	// 根据用户名，从数据库里获取用户，来判断是否某用户名是否已存在
	user, err := model.CheckUsername(username)
	if err != nil {
		log.Println("CheckUsername: ", err)
	}
	if user.ID > 0 {
		// 已存在用户名，不可用，返回msg
		w.Write([]byte("用户名已存在"))
	} else {
		// 用户名可用，返回msg
		w.Write([]byte("用户名可用"))
	}
}
