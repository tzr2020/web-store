package controller

import (
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"web-store/model"
	"web-store/util"
)

// registerUserRoutes 路由器注册用户相关的路由规则和处理器
func registerUserRoutes() {
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/regist", handlerRegist)
	http.HandleFunc("/checkUsername", handlerCheckUsername)
	http.HandleFunc("/logout", handlerLogout)
	http.HandleFunc("/uploadAvatar", uploadAvatar)
	http.HandleFunc("/user/space/index", userSpaceIndex)
	http.HandleFunc("/user/space/accountinfo", userSpaceAccountinfo)
}

// handlerLogin 用户登录处理器
func handlerLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// GET请求
		ok, _ := IsLogin(r)
		if ok {
			// 已经登录，返回首页，不用进行用户名和密码验证和创建Session操作
			w.Header().Set("Location", "/index")
			w.WriteHeader(302)
			return
		}

		// 返回登录页面
		getLoginPage(w, "")
	case http.MethodPost:
		// POST请求
		ok, _ := IsLogin(r)
		if ok {
			// 已经登录，返回首页，不用进行用户名和密码验证和创建Session操作
			w.Header().Set("Location", "/index")
			w.WriteHeader(302)
			return
		}

		// 从表单获取数据
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		// 从数据库获取用户信息
		user, err := model.CheckUsernameAndPassword(username, password)

		if err == model.ErrUserNotFound {
			// http.Error(w, err.Error(), http.StatusNotFound)
			getLoginPage(w, "用户名或密码错误")
			return
		}

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
			w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
			// w.Write([]byte(err.Error())) // 返回内部错误信息
			// http.Error(w, err.Error(), http.StatusInternalServerError) // 500 服务器错误 设置状态码和返回内部错误信息
			return
		}

		// 验证登录，返回页面
		CheckLogin(user, w)
	}
}

// getLoginPage 模板引擎生成最终页面，并返回登录页面
func getLoginPage(w http.ResponseWriter, msg string) {
	t, err := template.ParseFiles("./view/page/user/login.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}
	err = t.Execute(w, msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}
}

// CheckLogin 验证登录，返回最终页面
func CheckLogin(user *model.User, w http.ResponseWriter) {
	if user.ID > 0 {
		// 用户名和密码正确

		// 创建Session
		uuid := util.CreateUUID()
		sess := &model.Session{
			SessionID: uuid,
			Username:  user.Username,
			UserID:    user.ID,
		}

		// 将Session保存到数据库
		err := model.AddSession(sess)
		if err != nil {
			log.Println(err)
			http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
			return
		}

		// 创建Cookie
		cookie := http.Cookie{
			Name:   "user",
			Value:  sess.SessionID,
			MaxAge: 60 * 60 * 24 * 7,
			Path:   "/",
		}
		// 将Cookie发送到浏览器
		http.SetCookie(w, &cookie)

		// 登录成功，模板引擎生成最终页面，并返回首页
		w.Header().Set("Location", "/index")
		w.WriteHeader(302)
	} else {
		// 登录失败，返回登录页面
		getLoginPage(w, "用户名或密码错误")
	}
}

// handlerRegist 会员用户注册账号处理器
func handlerRegist(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET请求
	case http.MethodGet:
		// 返回注册页面
		getRegistPage(w, "")

	// POST请求
	case http.MethodPost:
		// 从表单获取用户注册信息
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")

		newUser := &model.User{
			Username: username,
			Password: password,
			Email:    email,
		}

		// 验证注册，返回页面
		CheckRegist(w, newUser)
	}
}

// getRegistPage 模板引擎生成最终页面，并返回注册页面
func getRegistPage(w http.ResponseWriter, msg string) {
	t, err := template.ParseFiles("./view/page/user/regist.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}
	err = t.Execute(w, msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}
}

// 验证注册信息，返回最终页面
func CheckRegist(w http.ResponseWriter, newUser *model.User) {
	// 从数据库查询是否存在某用户名的用户，来判断用户名是否可用
	user, err := model.CheckUsername(newUser.Username)
	if err != nil && err != model.ErrUserNotFound {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}

	if user.ID > 0 {
		// 用户名已存在，不可用
		// 返回注册页面，并带上提示信息
		getRegistPage(w, "用户名\""+newUser.Username+"\"已存在！")
	} else {
		// 用户名可用
		// 数据库新增用户
		err = model.AddUser(newUser)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
			w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
			return
		}

		// 返回登录页面
		w.Header().Set("Location", "/login")
		w.WriteHeader(302)
	}
}

// handlerCheckUsername 验证用户名处理器
func handlerCheckUsername(w http.ResponseWriter, r *http.Request) {
	// 从Ajax请求里获取用户输入的用户名
	username := r.PostFormValue("username")

	// 根据用户名，从数据库里获取用户，来判断是否某用户名是否已存在
	user, err := model.CheckUsername(username)
	if err != nil && err != model.ErrUserNotFound {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 500 服务器错误 设置状态码
		w.Write([]byte("服务器内部出现错误"))                  // 返回错误信息
		return
	}
	if user.ID > 0 {
		// 已存在用户名，不可用，返回msg
		w.Write([]byte("用户名已存在"))
	} else {
		// 用户名可用，返回msg
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}

// handlerLogout 会员用户注销账号处理器
func handlerLogout(w http.ResponseWriter, r *http.Request) {
	// 获取Cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		log.Printf("获取Cookie发生错误：%v", err)
		http.Error(w, "服务器内部发生错误", http.StatusInternalServerError)
		return
	}

	if cookie != nil {
		// 获取SessionID
		sessID := cookie.Value

		// 删除数据库中对应的Session
		err := model.DeleteSession(sessID)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return
		}

		// 设置Cookie为失效状态
		cookie.Path = "/"
		cookie.MaxAge = -1
		// 将Cookie发送到浏览器
		http.SetCookie(w, cookie)
	}

	// 注销账号后，返回登录页面
	w.Header().Set("Location", "/login")
	w.WriteHeader(302)
}

// IsLogin 根据Cookie，Session检查用户是否已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	// 获取Cookie
	cookie, _ := r.Cookie("user")

	if cookie != nil {
		// 获取SessionID
		sessID := cookie.Value
		// 获取Session
		sess, err := model.GetSession(sessID)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
		} else {
			if sess.UserID > 0 {
				// 用户已经登录
				return true, sess
			}
		}
	}

	return false, nil
}

// 上传用户头像
func uploadAvatar(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userID := r.PostFormValue("userID")

		// 解析请求包体，设置用于存储包体数据的内存大小
		r.ParseMultipartForm(32 << 20) // 32MB

		// 获取上传文件
		file, handler, err := r.FormFile("avatarFile")
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
				Msg:  "上传用户头像失败，图片文件后缀名要求是.jpg或.png",
			})
			return
		}

		// 存储上传文件本地
		path := "view/static/img/avatar/" + userID + suffix
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
		avatar := "/img/avatar/" + userID + ".jpg"
		err = model.UpdateUserAvatar(avatar, userID)
		if err != nil {
			log.Println(err)
			util.ResponseWriteJsonOfInsideServer(w)
			return
		}

		util.ResponseWriteJson(w, util.Json{
			Code: 200,
			Msg:  "上传用户头像成功",
		})
	}
}

// userSpaceIndex 返回会员用户的个人空间首页
func userSpaceIndex(w http.ResponseWriter, r *http.Request) {
	// 判断会员用户是否已经登录
	ok, _ := IsLogin(r)
	if !ok {
		// 重定向到会员用户登录页面
		w.Header().Set("Location", "/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的会员用户列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/user-space-layout.html",
		"./view/template/user-space-index.html",
	})
}

// userSpaceAccountinfo 返回会员用户的个人空间的账号信息页面
func userSpaceAccountinfo(w http.ResponseWriter, r *http.Request) {
	// 判断会员用户是否已经登录
	ok, _ := IsLogin(r)
	if !ok {
		// 重定向到会员用户登录页面
		w.Header().Set("Location", "/login")
		w.WriteHeader(302)
		return
	}
	// 返回后台管理系统的会员用户列表页面
	util.ExecuteTowTpl(w, [2]string{
		"./view/template/user-space-layout.html",
		"./view/template/user-space-accountinfo.html",
	})
}
