

# gin web 框架概览


### 核心接口 IRoutes

```go

// 注册路由的抽象
// IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...HandlerFunc) IRoutes  // 提供了用户接入自定义逻辑的能力，一般情况下也被看作是插件机制

	Handle(string, string, ...HandlerFunc) IRoutes

	// Any(string, ...HandlerFunc) IRoutes  // 对 handle 的二次包装，（见，范例 - 1）
	// GET(string, ...HandlerFunc) IRoutes
	// POST(string, ...HandlerFunc) IRoutes
	// DELETE(string, ...HandlerFunc) IRoutes
	// PATCH(string, ...HandlerFunc) IRoutes
	// PUT(string, ...HandlerFunc) IRoutes
	// OPTIONS(string, ...HandlerFunc) IRoutes
	// HEAD(string, ...HandlerFunc) IRoutes

	// StaticFile(string, string) IRoutes  // 额外提供了静态文件的接口，有平替（见，范例 - 2）
	// StaticFileFS(string, string, http.FileSystem) IRoutes
	// Static(string, string) IRoutes
	// StaticFS(string, http.FileSystem) IRoutes
}


// 函数式，作为参数，需符合签名
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)



// 范例 - 1

func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	if matched := regEnLetter.MatchString(httpMethod); !matched {
		panic("http method " + httpMethod + " is not valid")
	}
	return group.handle(httpMethod, relativePath, handlers)
}

// POST 是 router.Handle("POST", path, handle) 的快捷键
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodPost, relativePath, handlers)
}


// 范例 - 2

router.GET("/static", func(context *gin.Context) {
    // 读文件
    // 写响应
    // 总之，拎出个静态文件处理没啥必要
})

```




### Engine 实现

- 实现了路由树功能，提供了注册和匹配路由的功能      
    e.g. addRoute()


- 它本身可以作为一个 Handler 传递到 http 包，用于启动服务器。

Engine 的路由树功能本质上是依赖于 methodTree 的



### methodTrees 和 methodTree

    methodTree 才是真实的路由树
    Gin 定义了 methodTrees，它实际上代表的是森林，即每一个 HTTP 方法都对应到一棵树

```go
// Create an instance of Engine, by using New() or Default()
type Engine struct {
    // ...
	trees            methodTrees
}

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

```


### HandlerFunc 和 HandlersChain

    HandlerFunc 定义了核心抽象 -- 处理逻辑
    在默认情况下，它代表了注册路由的业务代码

    HandlersChain 则是构造了责任链模式

```go

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc 

    HandlerFunc1 -> HandlerFunc2 -> HandlerFunc3     
    PS. 最后一个则是封装了业务逻辑的 HandlerFunc 
```

### Context 抽象

Context 也是代表了执行的上下文，提供了丰富的 API：

- 处理请求的 API，代表的是以 Get 和 Bind 为前缀的方法 e.g. ctx.Get() ... 
- 处理响应的 API，例如返回 JSON 响应的方法 e.g. ctx.JSON()
- 渲染页面，如 HTML 方法

    ——— req ——— | ------- | ——— req ——>|
                | context |            | 业务逻辑
    <—— resp —— | ------- | ——— resp ——|

```go    

type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

    // ...
}

```



### 抽象总结

    methodTree 代表路由 & node 
    __________________

HandlerFunc 处理逻辑 / Context 上下文 / Engine 代表服务器






# 实现步骤

1. 定义 