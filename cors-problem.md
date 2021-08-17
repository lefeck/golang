> 当两个域具有相同的协议(如http), 相同的端口(如80)，相同的host，那么我们就可以认为它们是相同的域（协议，域名，端口都必须相同）。
> 跨域就指着**协议，域名，端口**不一致，出于安全考虑，跨域的资源之间是无法交互的(例如一般情况跨域的JavaScript无法交互，当然有很多解决跨域的方案)

**解决跨域几种方案**

```
/*
    CORS 普通跨域请求：只服务端设置Access-Control-Allow-Origin即可，
    前端无须设置，若要带cookie请求：前后端都需要设置。
		
    JSONP 缺点：只能使用get 请求
    document.domain 仅限主域相同，子域不同的跨域应用场景。
    window.name
    websockets
*/
```

**跨域资源共享CORS**

> CORS是一个W3C标准，全称是"跨域资源共享"（Cross-origin resource sharing）。
>
> 它允许浏览器向跨源服务器，发出[`XMLHttpRequest`](http://www.ruanyifeng.com/blog/2012/09/xmlhttprequest_level_2.html)请求，从而克服了AJAX只能[同源](http://www.ruanyifeng.com/blog/2016/04/same-origin-policy.html)使用的限制。

**Cors简介**

> CORS需要浏览器和服务器同时支持。目前，所有浏览器都支持该功能，IE浏览器不能低于IE10。
>
> 整个CORS通信过程，都是浏览器自动完成，不需要用户参与。对于开发者来说，CORS通信与同源的AJAX通信没有差别，代码完全一样。浏览器一旦发现AJAX请求跨源，就会自动添加一些附加的头信息，有时还会多出一次附加的请求，但用户不会有感觉。
>
> 因此，实现CORS通信的关键是服务器。只要服务器实现了CORS接口，就可以跨源通信。

**两种请求**

> 浏览器将CORS请求分成两类：简单请求（simple request）和非简单请求（not-so-simple request）。

满足以下两大条件,属于简单请求

* 请求方法是以下三种方法之一：
          HEAD
          GET
          POST
* HTTP的头信息不超出以下几种字段：
          Accept
          Accept-Language
          Content-Language
          Last-Event-ID
          Content-Type：只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain
  

> 这是为了兼容表单（form），因为历史上表单一直可以发出跨域请求。AJAX 的跨域设计就是，只要表单可以发，AJAX 就可以直接发。
>
> 凡是不同时满足上面两个条件，就属于非简单请求。
>
> 浏览器对这两种请求的处理，是不一样的。

**简单请求**

基本流程

> 对于简单请求，浏览器直接发出CORS请求。具体来说，就是在头信息之中，增加一个`Origin`字段。
>
> 下面是一个例子，浏览器发现这次跨源AJAX请求是简单请求，就自动在头信息之中，添加一个`Origin`字段。

```
/*
    GET /cors HTTP/1.1
    Origin: http://api.bob.com
    Host: api.alice.com
    Accept-Language: en-US
    Connection: keep-alive
    User-Agent: Mozilla/5.0...
*/
```

> 上面的头信息中，`Origin`字段用来说明，本次请求来自哪个源（协议 + 域名 + 端口）。服务器根据这个值，决定是否同意这次请求。
>
> 如果`Origin`指定的源，不在许可范围内，服务器会返回一个正常的HTTP回应。浏览器发现，这个回应的头信息没有包含`Access-Control-Allow-Origin`字段（详见下文），就知道出错了，从而抛出一个错误，被`XMLHttpRequest`的`onerror`回调函数捕获。注意，这种错误无法通过状态码识别，因为HTTP回应的状态码有可能是200。
>
> 如果`Origin`指定的域名在许可范围内，服务器返回的响应，会多出几个头信息字段。



```
/*
    Access-Control-Allow-Origin: http://api.bob.com
    Access-Control-Allow-Credentials: true
    Access-Control-Expose-Headers: FooBar
    Content-Type: text/html; charset=utf-8
*/
```

> 上面的头信息之中，有三个与CORS请求相关的字段，都以`Access-Control-`开头。

```
1 . Access-Control-Allow-Origin
```

> 该字段是必须的,他的值要么是请求Origin字段的值,要么是一个*, 表示接受任意域名的请求.

```
2 . Access-Control-Allow-Credentials
```

> 该字段可选。它的值是一个布尔值，表示是否允许发送Cookie。默认情况下，Cookie不包括在CORS请求之中。设为`true`，即表示服务器明确许可，Cookie可以包含在请求中，一起发给服务器。这个值也只能设为`true`，如果服务器不要浏览器发送Cookie，删除该字段即可.

```
3 . Access-Control-Expose-Headers
```

> 该字段可选。CORS请求时，`XMLHttpRequest`对象的`getResponseHeader()`方法只能拿到6个基本字段：`Cache-Control`、`Content-Language`、`Content-Type`、`Expires`、`Last-Modified`、`Pragma`。如果想拿到其他字段，就必须在`Access-Control-Expose-Headers`里面指定。上面的例子指定，`getResponseHeader('FooBar')`可以返回`FooBar`字段的值。

***3\***|***4\*****withCredentials属性**

> 上面说到，CORS请求默认不发送Cookie和HTTP认证信息。如果要把Cookie发到服务器，一方面要服务器同意，指定`Access-Control-Allow-Credentials`字段。



```
// Access-Control-Allow-Credentials: true
另一方面，开发者必须在AJAX请求中打开withCredentials属性。
```



```
// var xhr = new XMLHttpRequest();
// xhr.withCredentials = true;
```

> 否则，即使服务器同意发送Cookie，浏览器也不会发送。或者，服务器要求设置Cookie，浏览器也不会处理。
>
> 但是，如果省略`withCredentials`设置，有的浏览器还是会一起发送Cookie。这时，可以显式关闭`withCredentials`。



```
// xhr.withCredentials = false;
```

> 需要注意的是，如果要发送Cookie，`Access-Control-Allow-Origin`就不能设为星号，必须指定明确的、与请求网页一致的域名。同时，Cookie依然遵循同源政策，只有用服务器域名设置的Cookie才会上传，其他域名的Cookie并不会上传，且（跨源）原网页代码中的`document.cookie`也无法读取服务器域名下的Cookie。

***3\***|***5\*****非简单请求**

预检请求

> 非简单请求是那种对服务器有特殊要求的请求，比如请求方法是`PUT`或`DELETE`，或者`Content-Type`字段的类型是`application/json`。
>
> 非简单请求的CORS请求，会在正式通信之前，增加一次HTTP查询请求，称为"预检"请求（preflight）。
>
> 浏览器先询问服务器，当前网页所在的域名是否在服务器的许可名单之中，以及可以使用哪些HTTP动词和头信息字段。只有得到肯定答复，浏览器才会发出正式的`XMLHttpRequest`请求，否则就报错。
>
> 下面是一段浏览器的JavaScript脚本。

```
var url = 'http://api.alice.com/cors';
var xhr = new XMLHttpRequest();
xhr.open('PUT', url, true);
xhr.setRequestHeader('X-Custom-Header', 'value');
xhr.send();
```

> 上面代码中，HTTP请求的方法是`PUT`，并且发送一个自定义头信息`X-Custom-Header`。
>
> 浏览器发现，这是一个非简单请求，就自动发出一个"预检"请求，要求服务器确认可以这样请求。下面是这个"预检"请求的HTTP头信息。

```
OPTIONS /cors HTTP/1.1
Origin: http://api.bob.com
Access-Control-Request-Method: PUT
Access-Control-Request-Headers: X-Custom-Header
Host: api.alice.com
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```

> "预检"请求用的请求方法是`OPTIONS`，表示这个请求是用来询问的。头信息里面，关键字段是`Origin`，表示请求来自哪个源。

```
除了Origin字段，"预检"请求的头信息包括两个特殊字段
```



```
/*
  （1）Access-Control-Request-Method
  该字段是必须的，用来列出浏览器的CORS请求会用到哪些HTTP方法，上例是PUT。

  （2）Access-Control-Request-Headers
  该字段是一个逗号分隔的字符串，指定浏览器CORS请求会额外发送的头信息字段，上例是X-Custom-Header。
*/
预检请求的回应
```

> 服务器收到"预检"请求以后，检查了`Origin`、`Access-Control-Request-Method`和`Access-Control-Request-Headers`字段以后，确认允许跨源请求，就可以做出回应。

```
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:15:39 GMT
Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http://api.bob.com
Access-Control-Allow-Methods: GET, POST, PUT
Access-Control-Allow-Headers: X-Custom-Header
Content-Type: text/html; charset=utf-8
Content-Encoding: gzip
Content-Length: 0
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
Content-Type: text/plain
```

> 上面的HTTP回应中，关键的是`Access-Control-Allow-Origin`字段，表示`http://api.bob.com`可以请求数据。该字段也可以设为星号，表示同意任意跨源请求。

```
Access-Control-Allow-Origin: *
```

> 如果服务器否定了"预检"请求，会返回一个正常的HTTP回应，但是没有任何CORS相关的头信息字段。这时，浏览器就会认定，服务器不同意预检请求，因此触发一个错误，被`XMLHttpRequest`对象的`onerror`回调函数捕获。控制台会打印出如下的报错信息。

```
XMLHttpRequest cannot load http://api.alice.com.
Origin http://api.bob.com is not allowed by Access-Control-Allow-Origin.
服务器回应的其他CORS相关字段如下。
```



```
/*
    Access-Control-Allow-Methods: GET, POST, PUT
    Access-Control-Allow-Headers: X-Custom-Header
    Access-Control-Allow-Credentials: true
    Access-Control-Max-Age: 1728000
*/
```

**字段说明**

```
1.Access-Control-Allow-Origin
```

> 首先，客户端请求时要带上一个Origin，用来说明，本次请求来自哪个源（协议 + 域名 + 端口）。服务器根据这个值，决定是否同意这次请求。然后服务端在返回时需要带上这个字段，并把对方传过来的值返回去。告知客户端，允许这次请求。
> 这个字段也可以设置为*，即允许所有客户端访问。但是这样做会和Access-Control-Allow-Credentials 起冲突。可能导致跨域请求失败。

```
2.Access-Control-Allow-Credentials
```

> 这个字段是一个**BOOL**值，可以允许客户端携带一些校验信息，比如cookie等。如果设置为Access-Control-Allow-Origin：*，而该字段是true，并且客户端开启了withCredentials, 仍然不能正确访问。需要把Access-Control-Allow-Origin的值设置为客户端传过来的值。

```
3.Access-Control-Allow-Credentials
```

> 该字段与简单请求时的含义相同。

```
4.Access-Control-Max-Age
```

> 该字段可选，用来指定本次预检请求的有效期，单位为秒。上面结果中，有效期是20天（1728000秒），即允许缓存该条回应1728000秒（即20天），在此期间，不用发出另一条预检请求。

**浏览器的正常请求和回应**

> 一旦服务器通过了"预检"请求，以后每次浏览器正常的CORS请求，就都跟简单请求一样，会有一个`Origin`头信息字段。服务器的回应，也都会有一个`Access-Control-Allow-Origin`头信息字段。
>
> 下面是"预检"请求之后，浏览器的正常CORS请求。

```
/*
		PUT /cors HTTP/1.1
    Origin: http://api.bob.com
    Host: api.alice.com
    X-Custom-Header: value
    Accept-Language: en-US
    Connection: keep-alive
    User-Agent: Mozilla/5.0...
*/
```

> 上面头信息的`Origin`字段是浏览器自动添加的。
>
> 下面是服务器正常的回应。



```
/*
    Access-Control-Allow-Origin: http://api.bob.com
    Content-Type: text/html; charset=utf-8
*/
```

***3\***|***8\*****与JSONP的比较**



```
/*
		CORS与JSONP的使用目的相同，但是比JSONP更强大。

		JSONP只支持GET请求，CORS支持所有类型的HTTP请求。JSONP的优势在于支持老式浏览器，
		以及可以向不支持CORS的网站请求数据。
*/
```

***4\***|***0\*****开启中间件进行跨域*****4\***|***1\*****安装cors包**

```
govendor fetch github.com/gin-contrib/cors
```

**配置cors跨域**

**Example 1 :**

```go
package main
import (
    "github.com/gin-gonic/gin"
    "strings"
    "fmt"
    "net/http"
)

func main() {
        r := gin.Default()
        r.Use(Cors()) //开启中间件 允许使用跨域请求
        r.run()
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") //请求头部
        if origin != "" {
            //接收客户端发送的origin （重要！）
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin) 
            //服务器支持的所有跨域请求的方法
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") 
            //允许跨域设置可以返回其他子段，可以自定义字段
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
            // 允许浏览器（客户端）可以解析的头部 （重要）
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers") 
            //设置缓存时间
            c.Header("Access-Control-Max-Age", "172800") 
            //允许客户端传递校验信息比如 cookie (重要)
            c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                                                                                          
        }

        //允许类型校验 
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "ok!")
        }

        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic info is: %v", err)
            }
        }()

        c.Next()
    }
}

```

**Example 2 :**

```go
// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
 
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
 
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

```

**Example 3 :** 

```go
package middleware

import (
	"github.com/gin-contrib/cors" //import package
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		},
	)
}
```