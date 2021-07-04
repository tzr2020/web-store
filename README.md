# 农副产品销售网站

前台首页：[http://localhost:8081/index](http://localhost:8081/index)

后台首页：[http://localhost:8081/manage/index](http://localhost:8081/manage/index)

# 技术栈

## 后端技术栈

* 使用 Go 语言内置的 HTTP 服务器（net/http 包）作为 Web服务器
* 使用 Go 语言内置的模板引擎（html/template 包）来生成包含动态数据的 HTML 文档
* 选择 MySQL 作为数据库
## 前端技术栈

* 使用 jQuery 类库向后端发送异步请求获取动态数据
* 使用 Layui 框架设计后台管理页面
# 功能模块

## 前台模块

* 会员用户模块
    * 注册账号
    * 登录账号
    * 注销账号
* 产品展示模块
    * 所有产品（可根据产品类别，价格区间和页码来过滤产品）
    * 产品搜索
    * 产品详情
* 购物车模块
    * 产品添加到购物车
    * 查看购物车的购物项
    * 清空购物车的购物项
    * 删除购物车的购物项
    * 修改购物车的购物项数量
* 订单模块
    * 填写订单
    * 提交订单
    * 查看订单
    * 查看订单详情
    * 订单付款（通过更新订单的付款时间来模拟付款）
    * 确认收货
## 后台模块

* 管理员模块
    * 登录账号
    * 注销账号
* 会员用户模块
    * 用户列表
    * 用户地址列表
    * 用户 Session 列表
* 产品模块
    * 产品列表
    * 产品类别列表
* 购物车模块
    * 购物车列表
    * 购物项列表
* 订单模块
    * 订单列表
    * 订单支付类型列表
    * 订单状态字典列表
    * 订单地址列表
    * 订单项列表
* 首页栏目模块
    * 热卖产品列表
    * 最新产品列表
    * 推荐产品列表
    * 导航栏产品列表
# 采用 B/S 结构和 MVC 模式进行开发

在B/S结构中，B代表Browser即浏览器，S代表Server即服务器。用户通过使用浏览器（客户端），访问网站网页提供的操作界面，来与服务端进行交互，获取服务。两端通过HTTP协议进行数据传输。

在MVC开发模式中，M代表Model即模型，V代表View即视图，C代表Controller即控制器。根据分层原则，将Web开发中的代码解耦，使得系统具有可扩展性、易维护性和可重用性，方便程序员维护和阅读代码。

![image](https://raw.githubusercontent.com/tzr2020/Github-Article/main/Image/GoWebMVC开发模式.png)

Controller 层负责业务中的处理逻辑，主要流程有获取数据、处理数据和返回数据。

在获取数据步骤中，可以从HTTP请求的请求路径（path）、请求方法（method）、请求参数（params）等地方获取数据，也可从数据库获取数据。

在处理数据步骤中，一般有比较、计算和设置数据等操作，数据库的 CRUD 操作。

在返回数据步骤中，可以返回文件数据，如静态文件；可以返回由模板引擎（Template）生成的页面，如包含动态数据的 HTML 文档；可以返回进行序列化处理后的数据，如 JSON 数据。

Model 层负责对业务中的实体进行建模，操纵数据库。

View 层负责业务中的数据的展示。

Controller 层负责业务中的处理逻辑，主要流程：

1. 获取数据
    * 从 HTTP 请求获取
        * 请求路径（path）
        * 请求方法（method）
        * 请求参数（params）
            * 查询字符串参数
            * 表单参数
        * 请求头部（harders）
    * 从数据库获取
2. 处理数据
    * 比较、计算、设置数据
    * 操纵数据库（CRUD 操作）
3. 返回数据
    * 文件（如静态资源文件：CSS 样式文件，JS 脚本文件，图片文件）
    * 模版引擎（Template）执行模版生成的页面（如包含动态数据的 HTML 文档）
    * 序列化处理后的数据（如 JSON 数据）
    * 无（响应主体无数据）

Model 层负责对业务中的实体进行建模，操纵数据库。

View 层负责业务中的数据的展示。

网站一次服务的一般流程如下：

第一步：客户端（浏览器）发送HTTP请求。

第二部：服务端（Web应用程序）监听端口，接收HTTP请求并解析为结构体，根据自定义的路由规则，开启一个协程，将请求交给对应的处理器处理后返回HTTP响应。

第三步：客户端（浏览器）接收 HTTP 响应，根据响应内容的格式来渲染页面。

# 代码整体介绍

可配合我写的学习笔记往下观看。

* [Go HTTP服务](https://www.mubucm.com/doc/65nFd-NZppb)
* [Golang HTTP服务器实现原理](https://github.com/tzr2020/Github-Article/blob/main/Golang/Golang%20HTTP%E6%9C%8D%E5%8A%A1%E5%99%A8%E5%AE%9E%E7%8E%B0%E5%8E%9F%E7%90%86.md)
* [Golang Web应用中的模板引擎](https://github.com/tzr2020/Github-Article/blob/main/Golang/Golang%20Web%E5%BA%94%E7%94%A8%E4%B8%AD%E7%9A%84%E6%A8%A1%E6%9D%BF%E5%BC%95%E6%93%8E.md)
## main.go 程序入口

* 设置日志输出的标签
* 调用总体路由函数
* 配置并启动服务器
    * 使用 Go 语言内置的，默认的路由器
## util 包

### **db.go**

负责连接数据库，向 Model 层提供 *sql.DB 实例来操作数据库。

Go 语言的 database/sql 包只提供了操纵数据库的接口，而实现操纵数据库的逻辑需要我们导入第三方包，这些包实现了 Go 语言的操纵数据库的接口。

操纵 MySQL 第三方包：github.com/go-sql-driver/mysql

## controller 包

### 总体路由 router.go

使用 http.Handle 和 http.HandleFunc 函数为处理器注册路由。

**1、静态文件路由**

Go 语言提供了文件服务器的实现。

```go
http.StripPrefix("/static/", http.FileServer(http.Dir("view/static")))
http.StripPrefix("/manage/", http.FileServer(http.Dir("view/template/manage")))
```

**2、前台处理器路由**

调用各个前台模块的注册路由函数。

**3、后台处理器路由**

调用各个后台模块的注册路由函数。

**4、处理异步请求处理器路由**

调用处理异步请求的注册路由函数。

### 模块的处理器函数

处理器函数包含了业务的处理逻辑，而且处理器函数的形参规定为 http.ResponseWriter 和 http.Request 类型，分别代表响应和请求，可以从这两个形参变量获取请求的信息，返回响应。

## model 包

通过声明结构体来抽象模拟现实业务中的实体，编写方法来操作数据库。

结构体的字段一般都包含了数据库表的字段。

有时需要结构体实例序列化为 JSON 文本，通常需要为每个字段添加标签。

通过调用 Prepare 方法来预编译编写好的 SQL 语句，调用 Exec 方法来执行 SQL 语句。

调用 QueryRow 方法来执行 SQL 语句得到结果，再通过 Scan 方法来将结果扫描到模型实例。

调用 Query 方法来执行 SQL 语句得到结果集，再通过 Next 方法配合 for 语句来遍历结果集得到每一个结果，然后通过 Scan 方法来将结果扫描到模型实例，最后将每个模型实例添加到切片。

## view 包

模板文件和静态文件。


