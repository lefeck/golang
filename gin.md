# Gin框架介绍及使用

`Gin`是一个用Go语言编写的web框架。它是一个类似于`martini`但拥有更好性能的API框架, 由于使用了`httprouter`，速度提高了近40倍。 如果你是性能和高效的追求者, 你会爱上`Gin`。

## Gin框架介绍

Go世界里最流行的Web框架，[Github](https://github.com/gin-gonic/gin)上有`32K+`star。 基于[httprouter](https://github.com/julienschmidt/httprouter)开发的Web框架。 [中文文档](https://gin-gonic.com/zh-cn/docs/)齐全，简单易用的轻量级框架。

## Gin框架安装与使用

### 安装

下载并安装`Gin`:

```bash
go get -u github.com/gin-gonic/gin
```

### 第一个Gin示例：

我们需要先构建路由对象，也就是 `gin.Engine` 路由对象，之后在路由对象上注册请求路径对应的处理器（包括中间件），最后通过路由对象的 `router.Run` 方法启动监听。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行注册了请求处理器，就是后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
```

将上面的代码保存并编译执行，然后使用浏览器打开`127.0.0.1:8080/hello`就能看到一串JSON字符串。



### 创建Engine

在gin框架中，Engine被定义成为一个结构体，Engine代表gin框架的一个结构体定义，其中包含了路由组、中间件、页面渲染接口、框架配置设置等相关内容。默认的Engine可以通过gin.Default进行创建，或者使用gin.New()同样可以创建。两种方式如下所示：

engine1 = gin.Default() 

engine2 = gin.New()

gin.Default()和gin.New()的区别在于gin.Default也使用gin.New()创建engine实例，但是会默认使用Logger和Recovery中间件。

Logger是负责进行打印并输出日志的中间件，方便开发者进行程序调试；

Recovery中间件的作用是如果程序执行过程中遇到panic中断了服务，则Recovery会恢复程序执行，并返回服务器500内部错误。通常情况下，我们使用默认的gin.Default创建Engine实例。



### **处理HTTP请求**



**通用处理请求：**

> func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes

* httpMethod：第一个参数表示要处理的HTTP的请求类型，是GET/POST/DELETE等请求类型中的一种。
* relativePath：第二个参数表示要解析的接口，由开发想着进行定义。
* handlers：第三个参数是处理对应的请求的代码的定义。

```
engine.Handle("GET", "/hello", func(context *gin.Context) {
        fmt.Println(context.FullPath()) //获取请求路由
        userName := context.Query("name")
        fmt.Println(userName)
        context.Writer.Write([]byte("hello: " + userName))
    })
```

Context是gin框架中封装的一个结构体，这是gin框架中最重要，最基础的一个结构体对象。该结构体可以提供我们操作请求，处理请求，获取数据等相关的操作，通常称之为上下文对象，简单说为我们提供操作环境。

 

**分类处理请求：**

路由系统支持任意方式的请求，如下的方法用来提供对应方法来接收请求：

> func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes.  
> func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes.  



**GET 方法请求处理器**

可以通过context.Query和context.DefaultQuery获取GET请求携带的参数。

```
engine.GET("/hello", func(context *gin.Context) {
   username := context.Query("name")
   context.Writer.Write([]byte("hello world!" + username))
})
```

context.DefaultQuery： 除了context.DefaultQuery方法获取请求携带的参数数据以外，还可以使用context.Query方法来获取Get请求携带的参数。

**POST 方法请求处理器**

```
engine.POST("/login", func(context *gin.Context) {
   username, exist := context.GetPostForm("username")
   if exist {
      fmt.Println(username)
   }
   context.Writer.Write([]byte("hello world!" + username))
})
```

context.GetPostForm获取表单数据：POST请求以表单的形式提交数据,除了可以使用context.PostForm获取表单数据意外，还可以使用context.GetPostForm来获取表单数据。

**GET 方法请求处理器**

```
engine.DELETE("/user/:id", DeleteHandle)
func DeleteHandle(context *gin.Context) {
    userID := context.Param("id")
    context.Writer.Write([]byte("Delete user's id : " + userID))
}
```

客户端的请求接口是DELETE类型，请求url为：[http://localhost:8080/user/1](http://localhost:9000/user/1)。最后的1是要删除的用户的id，是一个变量。因此在服务端gin中，通过路由的:id来定义一个要删除用户的id变量值，同时使用context.Param进行获取。



**任意方法**

若需要监听同一个请求路径的多个任意方法，可以使用 Any()，这样就可以同时监听任意的请求方法了 ：

```
 func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes
```



### 处理器

handler 处理器，用于处理 HTTP 请求。在典型的 MVC 架构中也被叫做控制器的动作（action in controller）。是一个满足如下前面的函数：

```
type HandlerFunc func(*Context)
```

形式非常简单，就是一个可以接收 `*Context` 类型（`*gin.Context`）参数的函数即可。

处理器分为两类，中间件和请求处理器（业务逻辑处理器）。在 `router.GET()` 类的的方法中，最后一个 handler 就是请求处理器，除此之外前边的都是中间件，必须要有一个 handler 才可以。

处理器被调用时，会接收一个 `*Context` 参数，是请求上下文，我们用于获取请求和操作响应。请参考请求上下文，或请求，或响应章节，获得更多内容。



## RESTful API

REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

推荐阅读[阮一峰 理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们Web服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

| 请求方法 |     URL      |     含义     |
| :------: | :----------: | :----------: |
|   GET    |    /book     | 查询书籍信息 |
|   POST   | /create_book | 创建书籍记录 |
|   POST   | /update_book | 更新书籍信息 |
|   POST   | /delete_book | 删除书籍信息 |

同样的需求我们按照RESTful API设计如下：

| 请求方法 |  URL  |     含义     |
| :------: | :---: | :----------: |
|   GET    | /book | 查询书籍信息 |
|   POST   | /book | 创建书籍记录 |
|   PUT    | /book | 更新书籍信息 |
|  DELETE  | /book | 删除书籍信息 |

Gin框架支持开发RESTful API的开发。 string

```go
func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
}
```

开发RESTful API的时候我们通常使用[Postman](https://www.getpostman.com/)来作为客户端的测试工具。



## Gin渲染

### HTML渲染

我们首先定义一个存放模板文件的`templates`文件夹，然后在其内部按照业务分别定义一个`posts`文件夹和一个`users`文件夹。 `posts/index.html`文件的内容如下：

```template
{{define "posts/index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

`users/index.html`文件的内容如下：

```template
{{define "users/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>users/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

Gin框架中使用`LoadHTMLGlob()`或者`LoadHTMLFiles()`方法进行HTML模板渲染。

```go
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}
```

### 自定义模板函数

定义一个不转义相应内容的`safe`模板函数如下：

```go
func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://www.baidu.com'>百度</a>")
	})

	router.Run(":8080")
}
```

在`index.tmpl`中使用定义好的`safe`模板函数：

```template
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>修改模板引擎的标识符</title>
</head>
<body>
<div>{{ . | safe }}</div>
</body>
</html>
```

### 静态文件处理

当我们渲染的HTML文件中引用了静态文件时，我们只需要按照以下方式在渲染页面前调用`gin.Static`方法即可。

```go
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
   // ...
	r.Run(":8080")
}
```

### 使用模板继承

Gin框架默认都是使用单模板，如果需要使用`block template`功能，可以通过`"github.com/gin-contrib/multitemplate"`库实现，具体示例如下：

首先，假设我们项目目录下的templates文件夹下有以下模板文件，其中`home.tmpl`和`index.tmpl`继承了`base.tmpl`：

```bash
templates
├── includes
│   ├── home.tmpl
│   └── index.tmpl
├── layouts
│   └── base.tmpl
└── scripts.tmpl
```

然后我们定义一个`loadTemplates`函数如下：

```go
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
```

我们在`main`函数中

```go
func indexFunc(c *gin.Context){
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context){
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main(){
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}
```

### 补充文件路径处理

关于模板文件和静态文件的路径，我们需要根据公司/项目的要求进行设置。可以使用下面的函数获取当前执行程序的路径。

```go
func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}
```

### JSON渲染

```go
func main() {
	r := gin.Default()

	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}
```

### XML渲染

注意需要使用具名的结构体类型。

```go
func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}
```

### YMAL渲染

```go
r.GET("/someYAML", func(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
})
```

### protobuf渲染

```go
r.GET("/someProtoBuf", func(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	// protobuf 的具体定义写在 testdata/protoexample 文件中。
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	// 请注意，数据在响应中变为二进制数据
	// 将输出被 protoexample.Test protobuf 序列化了的数据
	c.ProtoBuf(http.StatusOK, data)
})
```

## 获取参数

### 获取querystring参数

`querystring`指的是URL中`?`后面携带的参数，例如：`/user/search?username=小王子&address=沙河`。 获取请求的querystring参数的方法如下：

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
```

### 获取form参数

请求的数据通过form表单来提交，例如向`/user/search`发送一个POST请求，获取请求数据的方式如下：

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}
```

### 获取path参数

请求的参数通过URL路径传递，例如：`/user/search/小王子/沙河`。 获取请求URL路径中的参数的方式如下。

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")
}
```



### 参数绑定

Gin框架提供给开发者表单实体绑定的功能，可以将表单数据与结构体绑定。

**表单实体绑定**

使用PostForm这种单个获取属性和字段的方式，代码量较多，需要一个一个属性进行获取。而表单数据的提交，往往对应着完整的数据结构体定义，其中对应着表单的输入项。gin框架提供了数据结构体和表单提交数据绑定的功能，提高表单数据获取的效率。

以一个用户注册功能来进行讲解表单实体绑定操作。用户注册需要提交表单数据，假设注册时表单数据包含三项，分别为：username、phone和password。

```go
type UserRegister struct {
    Username string form:"username" binding:"required"
    Phone    string form:"phone" binding:"required"
    Password string form:"password" binding:"required"
}
```

#### ShouldBindQuery

使用ShouldBindQuery可以实现Get方式的数据请求的绑定

```go
type User struct {
    UserName string `form:"name"`
    Classes  string `form:"classes"`
}

func main() {
    engine := gin.Default()

    // http://localhost:8080/hello?name=davie&classes=软件工程
    engine.GET("/hello", func(context *gin.Context) {
        fmt.Println(context.FullPath())
        user := User{}
        if err := context.ShouldBindQuery(&user); err != nil {
            fmt.Println("信息有误！")
        }
        fmt.Println(user.UserName)
        fmt.Println(user.Classes)
        context.Writer.Write([]byte("username: " + user.UserName + " classes: " + user.Classes))
    })

    engine.Run()
}
```

#### ShouldBind

使用ShouldBind可以实现Post方式的提交数据的绑定工作

为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的`Content-Type`识别请求数据类型并利用反射机制自动提取请求中`QueryString`、`form表单`、`JSON`、`XML`等参数到结构体中。 下面的示例代码演示了`.ShouldBind()`强大的功能，它能够基于请求自动提取`JSON`、`form表单`和`QueryString`类型的数据，并把值绑定到指定的结构体对象。

```go
// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
```

`ShouldBind`会按照下面的顺序解析请求中的数据完成绑定：

1. 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
2. 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。



#### ShouldBindJson

当客户端使用Json格式进行数据提交时，可以采用ShouldBindJson对数据进行绑定并自动解析

```go
func main() {
    engine := gin.Default()

    engine.POST("/addstudent", func(context *gin.Context) {
        fmt.Println(context.FullPath())
        var person Person
        if err := context.ShouldBindJSON(&person); err !=nil {
        //if err := context.BindJSON(&person); err != nil {
            fmt.Println("失败！")
        }
        fmt.Println(person.Name)
        fmt.Println(person.Sex)
        fmt.Println(person.Age)
        //context.Writer.Write([]byte(" 添加记录：" + person.Name))
        context.Writer.WriteString(context.FullPath())
    })

    engine.Run()
}

type Person struct {
    Name string `form:"name"`
    Sex  string `form:"sex"`
    Age  int    `form:"age"`
}
```



## 文件上传

### 单个文件上传

文件上传前端页面代码：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>上传文件示例</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="f1">
    <input type="submit" value="上传">
</form>
</body>
</html>
```

后端gin框架部分代码：

```go
func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	router.Run()
}
```

### 多个文件上传

```go
func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()
}
```

## 重定向

### HTTP重定向

HTTP 重定向很容易。 内部、外部重定向均支持。

```go
r.GET("/test", func(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
})
```

### 路由重定向

路由重定向，使用`HandleContext`：

```go
r.GET("/test", func(c *gin.Context) {
    // 指定重定向的URL
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
})
```

## Gin路由

### 普通路由

```go
r.GET("/index", func(c *gin.Context) {...})
r.GET("/login", func(c *gin.Context) {...})
r.POST("/login", func(c *gin.Context) {...})
```

此外，还有一个可以匹配所有请求方法的`Any`方法如下：

```go
r.Any("/test", func(c *gin.Context) {...})
```

为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回`views/404.html`页面。

```go
r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
```



### 路由组

gin框架中可以使用路由组来实现对路由的分类。

路由组是router.Group中的一个方法，用于对请求进行分组。可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对`{}`包裹同组的路由，用不用`{}`包裹功能上没什么区别。

```go
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {...})
		userGroup.GET("/login", func(c *gin.Context) {...})
		userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	r.Run()
}
```

路由组也是支持嵌套的，例如：

```go
shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {...})
	}
```

通常我们将路由分组用在划分业务逻辑或划分API版本时。

Group返回一个RouterGroup指针对象，而RouterGroup是gin框架中的一个路由组结构体定义。

```go
type RouterGroup struct {
    Handlers HandlersChain
    basePath string
    engine   *Engine
    root     bool
}
```

RouterGroup实现了IRoutes中定义的方法，包含统一处理请求的Handle和分类型处理的GET、POST等。

```go
type IRoutes interface {
    Use(...HandlerFunc) IRoutes

    Handle(string, string, ...HandlerFunc) IRoutes
    Any(string, ...HandlerFunc) IRoutes
    GET(string, ...HandlerFunc) IRoutes
    POST(string, ...HandlerFunc) IRoutes
    DELETE(string, ...HandlerFunc) IRoutes
    PATCH(string, ...HandlerFunc) IRoutes
    PUT(string, ...HandlerFunc) IRoutes
    OPTIONS(string, ...HandlerFunc) IRoutes
    HEAD(string, ...HandlerFunc) IRoutes

    StaticFile(string, string) IRoutes
    Static(string, string) IRoutes
    StaticFS(string, http.FileSystem) IRoutes
}
```

### 路由原理

Gin框架中的路由使用的是[httprouter](https://github.com/julienschmidt/httprouter)这个库。

其基本原理就是构造一个路由地址的前缀树。

## Gin中间件

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

### 定义中间件

Gin中的中间件必须是一个`gin.HandlerFunc`类型。例如我们像下面的代码一样定义一个统计请求耗时的中间件。

```go
// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
```

### 注册中间件

在gin框架中，我们可以为每个路由添加任意数量的中间件。

#### 为全局路由注册

```go
func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StatCost())
	
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
```

#### 为某个路由单独注册

```go
// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/test2", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
```

#### 为路由组注册中间件

为路由组注册中间件有以下两种写法。

写法1：

```go
shopGroup := r.Group("/shop", StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
```

写法2：

```go
shopGroup := r.Group("/shop")
shopGroup.Use(StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
```

### 中间件注意事项

#### gin默认中间件

`gin.Default()`默认使用了`Logger`和`Recovery`中间件，其中：

- `Logger`中间件将日志写入`gin.DefaultWriter`，即使配置了`GIN_MODE=release`。
- `Recovery`中间件会recover任何`panic`。如果有panic的话，会写入500响应码。

如果不想使用上面两个默认的中间件，可以使用`gin.New()`新建一个没有任何默认中间件的路由。

#### gin中间件中使用goroutine

当在中间件或`handler`中启动新的`goroutine`时，**不能使用**原始的上下文（c *gin.Context），必须使用其只读副本（`c.Copy()`）。



## 多数据格式返回请求结果

在gin框架中，支持返回多种请求数据格式

### []byte

```go
engine := gin.Default()
engine.GET("/hello", func(context *gin.Context) {
        fullPath := "请求路径：" + context.FullPath()
        fmt.Println(fullPath)
        context.Writer.Write([]byte(fullPath))
})
engine.Run()
```

使用context.Writer.Write向客户端写入返回数据。Writer是gin框架中封装的一个ResponseWriter接口类型，其中的write方法就是http.ResponseWriter中包含的方法。

```go
type ResponseWriter interface {
    http.ResponseWriter
    http.Hijacker
    http.Flusher
    http.CloseNotifier

    // Returns the HTTP response status code of the current request.
    Status() int

    // Returns the number of bytes already written into the response http body.
    // See Written()
    Size() int

    // Writes the string into the response body.
    WriteString(string) (int, error)

    // Returns true if the response body was already written.
    Written() bool

    // Forces to write the http header (status code + headers).
    WriteHeaderNow()

    // get the http.Pusher for server push
    Pusher() http.Pusher
}
```

### string

除了write方法以外，ResponseWriter自身还封装了WriteString方法返回数据

```go
// Writes the string into the response body.
WriteString(string) (int, error)
```

和[]byte类型调用一样，可以通过Writer进行调用

```go
engine.GET("/hello", func(context *gin.Context) {
        fullPath := "请求路径：" + context.FullPath()
        fmt.Println(fullPath)
        context.Writer.WriteString(fullPath)
})
```

### JSON

#### map类型

```go
engine := gin.Default()
engine.GET("/hellojson", func(context *gin.Context) {
    fullPath := "请求路径：" + context.FullPath()
    fmt.Println(fullPath)

    context.JSON(200, map[string]interface{}{
        "code":    1,
        "message": "OK",
        "data":    fullPath,
    })
})
engine.Run(":9000") 
```

调用JSON将map类型的数据转换成为json格式并返回给前端，第一个参数200表示设置请求返回的状态码。和http请求的状态码一致。

#### 结构体类型

除了map以外，结构体也是可以直接转换为JSON格式进行返回的

```go
//通用请求返回结构体定义
type Response struct {
    Code    int         json:"code"
    Message string      json:"msg"
    Data    interface{} json:"data"
}

engine.GET("/jsonstruct", func(context *gin.Context) {
    fullPath := "请求路径：" + context.FullPath()
    fmt.Println(fullPath)
    resp := Response{Code: 1, Message: "Ok", Data: fullPath}
    context.JSON(200, &resp)
})
```

### HTML模板

除了JSON格式以外，gin框架还支持返回HTML格式的数据

```go
engine := gin.Default()
//设置html的目录
engine.LoadHTMLGlob("./html/*")
engine.GET("/hellohtml", func(context *gin.Context) {
    fullPath := "请求路径:" + context.FullPath()

    context.HTML(http.StatusOK, "index.html", gin.H{
        "title":    "Gin框架",
        "fullpath": fullPath,
    })
})
engine.Run()
```

### 加载静态资源文件

如果需要再页面是添加一张img，需要将img所在的目录进行静态资源路径设置才可能会生效：

```
engine.Static("/img", "./img")
```

同理，在项目开发时，一些静态的资源文件如html、js、css等可以通过静态资源文件设置的方式来进行设置



## 运行多个服务

我们可以在多个端口启动服务，例如：

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
   // 借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
```





## gin的生命周期

看完gin框架流程我有大致如下几个感触：

- gin是我目前看过的这三个go框架里最简洁的框架
- gin和iris在框架设计存在风格一致的地方，例如注册中间件、handle的执行

下图就是我对整个Gin框架生命周期的输出，由于图片过大存在平台压缩的可能，建议大家直接查看原图链接。

![img](http://cdn.tigerb.cn/20190704211526.png)

## 关键代码解析

```go
// 获取一个gin框架实例
gin.Default()
⬇️
// 具体的Default方法
func Default() *Engine {
    // 调试模式日志输出 
    debugPrintWARNINGDefault()
    // 创建一个gin框架实例
    engine := New()
    // 注册中间件
    engine.Use(Logger(), Recovery())
    return engine
}
⬇️
// 创建一个gin框架实例 具体方法
func New() *Engine {
    // 调试模式日志输出 
    debugPrintWARNINGNew()
    // 先插入一个小话题，可能好多人都在想为什么叫gin呢？
    // 哈哈，这个框架实例的结构体实际命名的Engine, 很明显gin就是一个很个性的简称了，是不是真相大白了。
    // 初始化一个Engine实例
    engine := &Engine{
        // 路由组
        // 给框架实例绑定上一个路由组
        RouterGroup: RouterGroup{
            // engine.Use 注册的中间方法到这里
            Handlers: nil,
            basePath: "/",
            // 是否是路由根节点
            root:     true,
        },
        FuncMap:                template.FuncMap{},
        RedirectTrailingSlash:  true,
        RedirectFixedPath:      false,
        HandleMethodNotAllowed: false,
        ForwardedByClientIP:    true,
        AppEngine:              defaultAppEngine,
        UseRawPath:             false,
        UnescapePathValues:     true,
        MaxMultipartMemory:     defaultMultipartMemory,
        // 路由树
        // 我们的路由最终注册到了这里
        trees:                  make(methodTrees, 0, 9),
        delims:                 render.Delims{Left: "{{", Right: "}}"},
        secureJsonPrefix:       "while(1);",
    }
    // RouterGroup绑定engine自身的实例
    // 不太明白为何如此设计
    // 职责分明么？
    engine.RouterGroup.engine = engine
    // 绑定从实例池获取上下文的闭包方法
    engine.pool.New = func() interface{} {
        // 获取一个Context实例
        return engine.allocateContext()
    }
    // 返回框架实例
    return engine
}
⬇️
// 注册日志&goroutin panic捕获中间件
engine.Use(Logger(), Recovery())
⬇️
// 具体的注册中间件的方法
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
    engine.RouterGroup.Use(middleware...)
    engine.rebuild404Handlers()
    engine.rebuild405Handlers()
    return engine
}

// 上面 是一个engine框架实例初始化的关键代码
// 我们基本看完了
// --------------router--------------
// 接下来 开始看路由注册部分

// 注册GET请求路由
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
    // 往路由组内 注册GET请求路由
    return group.handle("GET", relativePath, handlers)
}
⬇️
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
    absolutePath := group.calculateAbsolutePath(relativePath)
    // 把中间件的handle和该路由的handle合并
    handlers = group.combineHandlers(handlers)
    // 注册一个GET集合的路由
    group.engine.addRoute(httpMethod, absolutePath, handlers)
    return group.returnObj()
}
⬇️
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
    assert1(path[0] == '/', "path must begin with '/'")
    assert1(method != "", "HTTP method can not be empty")
    assert1(len(handlers) > 0, "there must be at least one handler")

    debugPrintRoute(method, path, handlers)
    // 检查有没有对应method集合的路由
    root := engine.trees.get(method)
    if root == nil {
        // 没有 创建一个新的路由节点
        root = new(node)
        // 添加该method的路由tree到当前的路由到路由树里
        engine.trees = append(engine.trees, methodTree{method: method, root: root})
    }
    // 添加路由
    root.addRoute(path, handlers)
}
⬇️
// 很关键
// 路由树节点
type node struct {
    // 路由path
    path      string
    indices   string
    // 子路由节点
    children  []*node
    // 所有的handle 构成一个链
    handlers  HandlersChain
    priority  uint32
    nType     nodeType
    maxParams uint8
    wildChild bool
}
// --------------http server--------------
// 接下来 开始看gin如何启动的http server

func (engine *Engine) Run(addr ...string) (err error) {
    defer func() { debugPrintError(err) }()

    address := resolveAddress(addr)
    debugPrint("Listening and serving HTTP on %s\n", address)
    // 执行http包的ListenAndServe方法 启动路由
    // engine实现了http.Handler接口 所以在这里作为参数传参进去
    // 后面我们再看engine.ServeHTTP的具体逻辑
    err = http.ListenAndServe(address, engine)
    return
}
⬇️
// engine自身就实现了Handler接口
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
⬇️
// 下面就是网络相关
// 监听IP+端口
ln, err := net.Listen("tcp", addr)
⬇️
// 上面执行完了监听
// 接着就是Serve
srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
⬇️
// Accept请求
rw, e := l.Accept()
⬇️
// 使用goroutine去处理一个请求
// 最终就执行的是engine的ServeHTTP方法
go c.serve(ctx)

// 上面服务已经启动起来了
// --------------handle request--------------
// 接着我们来看看engine的ServeHTTP方法的具体内容
// engine实现http.Handler接口ServeHTTP的具体方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // 获取一个上下文实例
    // 从实例池获取 性能高
    c := engine.pool.Get().(*Context)
    // 重置获取到的上下文实例的http.ResponseWriter
    c.writermem.reset(w)
    // 重置获取到的上下文实例*http.Request
    c.Request = req
    // 重置获取到的上下文实例的其他属性
    c.reset()

    // 实际处理请求的地方
    // 传递当前的上下文
    engine.handleHTTPRequest(c)

    //归还上下文实例
    engine.pool.Put(c)
}
⬇️
// 具体执行路由的方法
engine.handleHTTPRequest(c)
⬇️
t := engine.trees
for i, tl := 0, len(t); i < tl; i++ {
    // 这里寻找当前请求method的路由树节点
    // 我在想这里为啥不用map呢？
    // 虽说也遍历不了几次
    if t[i].method != httpMethod {
        continue
    }
    // 找到节点
    root := t[i].root
    // 很关键的地方
    // 寻找当前请求的路由
    handlers, params, tsr := root.getValue(path, c.Params, unescape)
    if handlers != nil {
        // 把找到的handles赋值给上下文
        c.handlers = handlers
        // 把找到的入参赋值给上下文
        c.Params = params
        // 执行handle
        c.Next()
        // 处理响应内容
        c.writermem.WriteHeaderNow()
        return
    }
    ...
}

// 方法树结构体
type methodTree struct {
    // HTTP Method
    method string
    // 当前HTTP Method的路由节点
    root   *node
}

// 方法树集合
type methodTrees []methodTree
⬇️
// 执行handle
func (c *Context) Next() {
    // 上下文处理之后c.index被执为-1
    c.index++
    for s := int8(len(c.handlers)); c.index < s; c.index++ {
        // 遍历执行所有handle(其实就是中间件+路由handle)
        // 首先感觉这里的设计又是似曾相识 iris不是也是这样么 不懂了 哈哈
        // 其次感觉这里设计的很一般 遍历？多无聊，这里多么适合「责任链模式」
        // 之后给大家带来关于这个handle执行的「责任链模式」的设计
        c.handlers[c.index](c)
    }
}

// Context的重置方法
func (c *Context) reset() {
    c.Writer = &c.writermem
    c.Params = c.Params[0:0]
    c.handlers = nil
    // 很关键 注意这里是-1哦
    c.index = -1
    c.Keys = nil
    c.Errors = c.Errors[0:0]
    c.Accepted = nil
}
```

# gin框架源码解析

## gin框架路由详解

gin框架使用的是定制版本的[httprouter](https://github.com/julienschmidt/httprouter)，其路由的原理是大量使用公共前缀的树结构，它基本上是一个紧凑的[Trie tree](https://baike.sogou.com/v66237892.htm)（或者只是[Radix Tree](https://baike.sogou.com/v73626121.htm)）。具有公共前缀的节点也共享一个公共父节点。

### Radix Tree

基数树（Radix Tree）又称为PAT位树（Patricia Trie or crit bit tree），是一种更节省空间的前缀树（Trie Tree）。对于基数树的每个节点，如果该节点是唯一的子树的话，就和父节点合并。下图为一个基数树示例：

![img](https://www.liwenzhou.com/images/Go/gin/radix_tree.png)

`Radix Tree`可以被认为是一棵简洁版的前缀树。我们注册路由的过程就是构造前缀树的过程，具有公共前缀的节点也共享一个公共父节点。假设我们现在注册有以下路由信息：

```go
r := gin.Default()

r.GET("/", func1)
r.GET("/search/", func2)
r.GET("/support/", func3)
r.GET("/blog/", func4)
r.GET("/blog/:post/", func5)
r.GET("/about-us/", func6)
r.GET("/about-us/team/", func7)
r.GET("/contact/", func8)
```

那么我们会得到一个`GET`方法对应的路由树，具体结构如下：

```bash
Priority   Path             Handle
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>
```

上面最右边那一列每个`*<数字>`表示Handle处理函数的内存地址(一个指针)。从根节点遍历到叶子节点我们就能得到完整的路由表。

例如：`blog/:post`其中`:post`只是实际文章名称的占位符(参数)。与`hash-maps`不同，这种树结构还允许我们使用像`:post`参数这种动态部分，因为我们实际上是根据路由模式进行匹配，而不仅仅是比较哈希值。

由于URL路径具有层次结构，并且只使用有限的一组字符(字节值)，所以很可能有许多常见的前缀。这使我们可以很容易地将路由简化为更小的问题。此外，**路由器为每种请求方法管理一棵单独的树**。一方面，它比在每个节点中都保存一个method-> handle map更加节省空间，它还使我们甚至可以在开始在前缀树中查找之前大大减少路由问题。

为了获得更好的可伸缩性，每个树级别上的子节点都按`Priority(优先级)`排序，其中优先级（最左列）就是在子节点(子节点、子子节点等等)中注册的句柄的数量。这样做有两个好处:

1. 首先优先匹配被大多数路由路径包含的节点。这样可以让尽可能多的路由快速被定位。
2. 类似于成本补偿。最长的路径可以被优先匹配，补偿体现在最长的路径需要花费更长的时间来定位，如果最长路径的节点能被优先匹配（即每次拿子节点都命中），那么路由匹配所花的时间不一定比短路径的路由长。下面展示了节点（每个`-`可以看做一个节点）匹配的路径：从左到右，从上到下。

```bash
   ├------------
   ├---------
   ├-----
   ├----
   ├--
   ├--
   └-
```

### 路由树节点

路由树是由一个个节点构成的，gin框架路由树的节点由`node`结构体表示，它有以下字段：

```go
// tree.go

type node struct {
   // 节点路径，比如上面的s，earch，和upport
	path      string
	// 和children字段对应, 保存的是分裂的分支的第一个字符
	// 例如search和support, 那么s节点的indices对应的"eu"
	// 代表有两个分支, 分支的首字母分别是e和u
	indices   string
	// 儿子节点
	children  []*node
	// 处理函数链条（切片）
	handlers  HandlersChain
	// 优先级，子节点、子子节点等注册的handler数量
	priority  uint32
	// 节点类型，包括static, root, param, catchAll
	// static: 静态节点（默认），比如上面的s，earch等节点
	// root: 树的根节点
	// catchAll: 有*匹配的节点
	// param: 参数节点
	nType     nodeType
	// 路径上最大参数个数
	maxParams uint8
	// 节点是否是参数节点，比如上面的:post
	wildChild bool
	// 完整路径
	fullPath  string
}
```

### 请求方法树

在gin的路由中，每一个`HTTP Method`(GET、POST、PUT、DELETE…)都对应了一棵 `radix tree`，我们注册路由的时候会调用下面的`addRoute`函数：

```go
// gin.go
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
   // liwenzhou.com...
   
   // 获取请求方法对应的树
	root := engine.trees.get(method)
	if root == nil {
	
	   // 如果没有就创建一个
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}
```

从上面的代码中我们可以看到在注册路由的时候都是先根据请求方法获取对应的树，也就是gin框架会为每一个请求方法创建一棵对应的树。只不过需要注意到一个细节是gin框架中保存请求方法对应树关系并不是使用的map而是使用的切片，`engine.trees`的类型是`methodTrees`，其定义如下：

```go
type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree  // slice
```

而获取请求方法对应树的get方法定义如下：

```go
func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}
```

为什么使用切片而不是map来存储`请求方法->树`的结构呢？我猜是出于节省内存的考虑吧，毕竟HTTP请求方法的数量是固定的，而且常用的就那几种，所以即使使用切片存储查询起来效率也足够了。顺着这个思路，我们可以看一下gin框架中`engine`的初始化方法中，确实对`tress`字段做了一次内存申请：

```go
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		// liwenzhou.com ...
		// 初始化容量为9的切片（HTTP1.1请求方法共9种）
		trees:                  make(methodTrees, 0, 9),
		// liwenzhou.com...
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}
```

### 注册路由

注册路由的逻辑主要有`addRoute`函数和`insertChild`方法。

#### addRoute

```go
// tree.go

// addRoute 将具有给定句柄的节点添加到路径中。
// 不是并发安全的
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++
	numParams := countParams(path)  // 数一下参数个数

	// 空树就直接插入当前节点
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		// 更新当前节点的最大参数个数
		if numParams > n.maxParams {
			n.maxParams = numParams
		}

		// 找到最长的通用前缀
		// 这也意味着公共前缀不包含“:”"或“*” /
		// 因为现有键不能包含这些字符。
		i := longestCommonPrefix(path, n.path)

		// 分裂边缘（此处分裂的是当前树节点）
		// 例如一开始path是search，新加入support，s是他们通用的最长前缀部分
		// 那么会将s拿出来作为parent节点，增加earch和upport作为child节点
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],  // 公共前缀后的部分作为子节点
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1, //子节点优先级-1
				fullPath:  n.fullPath,
			}

			// Update maxParams (max of all children)
			for _, v := range child.children {
				if v.maxParams > child.maxParams {
					child.maxParams = v.maxParams
				}
			}

			n.children = []*node{&child}
			// []byte for proper unicode char conversion, see #65
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		// 将新来的节点插入新的parent节点作为子节点
		if i < len(path) {
			path = path[i:]

			if n.wildChild {  // 如果是参数节点
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++

				// Update maxParams of the child node
				if numParams > n.maxParams {
					n.maxParams = numParams
				}
				numParams--

				// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
					// 检查更长的通配符, 例如 :name and :names
					if len(n.path) >= len(path) || path[len(n.path)] == '/' {
						continue walk
					}
				}

				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(path, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}
			// 取path首字母，用来与indices做比较
			c := path[0]

			// 处理参数后加斜线情况
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			// 检查路path下一个字节的子节点是否存在
			// 比如s的子节点现在是earch和upport，indices为eu
			// 如果新加一个路由为super，那么就是和upport有匹配的部分u，将继续分列现在的upport节点
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			// 否则就插入
			if c != ':' && c != '*' {
				// []byte for proper unicode char conversion, see #65
				// 注意这里是直接拼接第一个字符到n.indices
				n.indices += string([]byte{c})
				child := &node{
					maxParams: numParams,
					fullPath:  fullPath,
				}
				// 追加子节点
				n.children = append(n.children, child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			}
			n.insertChild(numParams, path, fullPath, handlers)
			return
		}

		// 已经注册过的节点
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		return
	}
}
```

其实上面的代码很好理解，大家可以参照动画尝试将以下情形代入上面的代码逻辑，体味整个路由树构造的详细过程：

1. 第一次注册路由，例如注册search
2. 继续注册一条没有公共前缀的路由，例如blog
3. 注册一条与先前注册的路由有公共前缀的路由，例如support

![addroute](https://www.liwenzhou.com/images/Go/gin/addroute.gif)

#### insertChild

```go
// tree.go
func (n *node) insertChild(numParams uint8, path string, fullPath string, handlers HandlersChain) {
  // 找到所有的参数
	for numParams > 0 {
		// 查找前缀直到第一个通配符
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // 没有发现通配符
			break
		}

		// 通配符的名称必须包含':' 和 '*'
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// 检查通配符是否有名称
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		// 检查这个节点是否有已经存在的子节点
		// 如果我们在这里插入通配符，这些子节点将无法访问
		if len(n.children) > 0 {
			panic("wildcard segment '" + wildcard +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		if wildcard[0] == ':' { // param
			if i > 0 {
				// 在当前通配符之前插入前缀
				n.path = path[:i]
				path = path[i:]
			}

			n.wildChild = true
			child := &node{
				nType:     param,
				path:      wildcard,
				maxParams: numParams,
				fullPath:  fullPath,
			}
			n.children = []*node{child}
			n = child
			n.priority++
			numParams--

			// 如果路径没有以通配符结束
			// 那么将有另一个以'/'开始的非通配符子路径。
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]

				child := &node{
					maxParams: numParams,
					priority:  1,
					fullPath:  fullPath,
				}
				n.children = []*node{child}
				n = child  // 继续下一轮循环
				continue
			}

			// 否则我们就完成了。将处理函数插入新叶子中
			n.handlers = handlers
			return
		}

		// catchAll
		if i+len(wildcard) != len(path) || numParams > 1 {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
		}

		// currently fixed width 1 for '/'
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]
		
		// 第一个节点:路径为空的catchAll节点
		child := &node{
			wildChild: true,
			nType:     catchAll,
			maxParams: 1,
			fullPath:  fullPath,
		}
		// 更新父节点的maxParams
		if n.maxParams < 1 {
			n.maxParams = 1
		}
		n.children = []*node{child}
		n.indices = string('/')
		n = child
		n.priority++

		// 第二个节点:保存变量的节点
		child = &node{
			path:      path[i:],
			nType:     catchAll,
			maxParams: 1,
			handlers:  handlers,
			priority:  1,
			fullPath:  fullPath,
		}
		n.children = []*node{child}

		return
	}

	// 如果没有找到通配符，只需插入路径和句柄
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}
```

`insertChild`函数是根据`path`本身进行分割，将`/`分开的部分分别作为节点保存，形成一棵树结构。参数匹配中的`:`和`*`的区别是，前者是匹配一个字段而后者是匹配后面所有的路径。

### 路由匹配

我们先来看gin框架处理请求的入口函数`ServeHTTP`：

```go
// gin.go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  // 这里使用了对象池
	c := engine.pool.Get().(*Context)
  // 这里有一个细节就是Get对象后做初始化
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)  // 我们要找的处理HTTP请求的函数

	engine.pool.Put(c)  // 处理完请求后将对象放回池子
}
```

函数很长，这里省略了部分代码，只保留相关逻辑代码：

```go
// gin.go
func (engine *Engine) handleHTTPRequest(c *Context) {
	// liwenzhou.com...

	// 根据请求方法找到对应的路由树
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// 在路由树中根据path查找
		value := root.getValue(rPath, c.Params, unescape)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.Params = value.params
			c.fullPath = value.fullPath
			c.Next()  // 执行函数链条
			c.writermem.WriteHeaderNow()
			return
		}
	
	// liwenzhou.com...
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}
```

路由匹配是由节点的 `getValue`方法实现的。`getValue`根据给定的路径(键)返回`nodeValue`值，保存注册的处理函数和匹配到的路径参数数据。

如果找不到任何处理函数，则会尝试TSR(尾随斜杠重定向)。

代码虽然很长，但还算比较工整。大家可以借助注释看一下路由查找及参数匹配的逻辑。

```go
// tree.go

type nodeValue struct {
	handlers HandlersChain
	params   Params  // []Param
	tsr      bool
	fullPath string
}

// liwenzhou.com...

func (n *node) getValue(path string, po Params, unescape bool) (value nodeValue) {
	value.params = po
walk: // Outer loop for walking the tree
	for {
		prefix := n.path
		if path == prefix {
			// 我们应该已经到达包含处理函数的节点。
			// 检查该节点是否注册有处理函数
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			// 没有找到处理函数 检查这个路径末尾+/ 是否存在注册函数
			indices := n.indices
			for i, max := 0, len(indices); i < max; i++ {
				if indices[i] == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

		if len(path) > len(prefix) && path[:len(prefix)] == prefix {
			path = path[len(prefix):]
			// 如果该节点没有通配符(param或catchAll)子节点
			// 我们可以继续查找下一个子节点
			if !n.wildChild {
				c := path[0]
				indices := n.indices
				for i, max := 0, len(indices); i < max; i++ {
					if c == indices[i] {
						n = n.children[i] // 遍历树
						continue walk
					}
				}

				// 没找到
				// 如果存在一个相同的URL但没有末尾/的叶子节点
				// 我们可以建议重定向到那里
				value.tsr = path == "/" && n.handlers != nil
				return
			}

			// 根据节点类型处理通配符子节点
			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				end := 0
				for end < len(path) && path[end] != '/' {
					end++
				}

				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[1:]
				val := path[:end]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(val); err != nil {
						value.params[i].Value = val // fallback, in case of error
					}
				} else {
					value.params[i].Value = val
				}

				// 继续向下查询
				if end < len(path) {
					if len(n.children) > 0 {
						path = path[end:]
						n = n.children[0]
						continue walk
					}

					// ... but we can't
					value.tsr = len(path) == end+1
					return
				}

				if value.handlers = n.handlers; value.handlers != nil {
					value.fullPath = n.fullPath
					return
				}
				if len(n.children) == 1 {
					// 没有找到处理函数. 检查此路径末尾加/的路由是否存在注册函数
					// 用于 TSR 推荐
					n = n.children[0]
					value.tsr = n.path == "/" && n.handlers != nil
				}
				return

			case catchAll:
				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[2:]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(path); err != nil {
						value.params[i].Value = path // fallback, in case of error
					}
				} else {
					value.params[i].Value = path
				}

				value.handlers = n.handlers
				value.fullPath = n.fullPath
				return

			default:
				panic("invalid node type")
			}
		}

		// 找不到，如果存在一个在当前路径最后添加/的路由
		// 我们会建议重定向到那里
		value.tsr = (path == "/") ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)
		return
	}
}
```

## gin框架中间件详解

gin框架涉及中间件相关有4个常用的方法，它们分别是`c.Next()`、`c.Abort()`、`c.Set()`、`c.Get()`。

### 中间件的注册

gin框架中的中间件设计很巧妙，我们可以首先从我们最常用的`r := gin.Default()`的`Default`函数开始看，它内部构造一个新的`engine`之后就通过`Use()`函数注册了`Logger`中间件和`Recovery`中间件：

```go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())  // 默认注册的两个中间件
	return engine
}
```

继续往下查看一下`Use()`函数的代码：

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)  // 实际上还是调用的RouterGroup的Use函数
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}
```

从下方的代码可以看出，注册中间件其实就是将中间件函数追加到`group.Handlers`中：

```go
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}
```

而我们注册路由时会将对应路由的函数和之前的中间件函数结合到一起：

```go
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)  // 将处理请求的函数与中间件函数结合
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}
```

其中结合操作的函数内容如下，注意观察这里是如何实现拼接两个切片得到一个新切片的。

```go
const abortIndex int8 = math.MaxInt8 / 2

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {  // 这里有一个最大限制
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}
```

也就是说，我们会将一个路由的中间件函数和处理函数结合到一起组成一条处理函数链条`HandlersChain`，而它本质上就是一个由`HandlerFunc`组成的切片：

```go
type HandlersChain []HandlerFunc
```

### 中间件的执行

我们在上面路由匹配的时候见过如下逻辑：

```go
value := root.getValue(rPath, c.Params, unescape)
if value.handlers != nil {
  c.handlers = value.handlers
  c.Params = value.params
  c.fullPath = value.fullPath
  c.Next()  // 执行函数链条
  c.writermem.WriteHeaderNow()
  return
}
```

其中`c.Next()`就是很关键的一步，它的代码很简单：

```go
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}
```

从上面的代码可以看到，这里通过索引遍历`HandlersChain`链条，从而实现依次调用该路由的每一个函数（中间件或处理请求的函数）。

![gin_middleware1](https://www.liwenzhou.com/images/Go/gin/gin_middleware1.png)

我们可以在中间件函数中通过再次调用`c.Next()`实现嵌套调用（func1中调用func2；func2中调用func3），

![gin_middleware2](https://www.liwenzhou.com/images/Go/gin/gin_middleware2.png)

或者通过调用`c.Abort()`中断整个调用链条，从当前函数返回。

```go
func (c *Context) Abort() {
	c.index = abortIndex  // 直接将索引置为最大限制值，从而退出循环
}
```

### c.Set()/c.Get()

`c.Set()`和`c.Get()`这两个方法多用于在多个函数之间通过`c`传递数据的，比如我们可以在认证中间件中获取当前请求的相关信息（userID等）通过`c.Set()`存入`c`，然后在后续处理业务逻辑的函数中通过`c.Get()`来获取当前请求的用户。`c`就像是一根绳子，将该次请求相关的所有的函数都串起来了。![gin_middleware3](https://www.liwenzhou.com/images/Go/gin/gin_middleware3.png)

## 总结

1. gin框架路由使用前缀树，路由注册的过程是构造前缀树的过程，路由匹配的过程就是查找前缀树的过程。
2. gin框架的中间件函数和处理函数是以切片形式的调用链条存在的，我们可以顺序调用也可以借助`c.Next()`方法实现嵌套调用。
3. 借助`c.Set()`和`c.Get()`方法我们能够在不同的中间件函数中传递数据。


