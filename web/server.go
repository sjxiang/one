package web

import (
	"net"
	"net/http"
)

// 评论：采用动词 Handle 更符合 Go 的命名风格
type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	// Start 启动服务器
	// addr 是监听地址，如果只指定端口，可以使用 “:8081”，或者 “localhost:8081”
	Start(addr string) error 
	// AddRoute 注册一个路由
	// method 是 HTTP 方法
	// path 是路径，必须以 / 开头
	AddRoute(method, path string, handler HandleFunc)
}


// 确保一定实现了 Server 接口
var _ Server = &HTTPServer{}


type HTTPServer struct {

}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

// 评论：言简意赅，不像 gin 核心接口 IRoutes 中的 Handle 模棱两可，看上去像是处理什么东西，而实质上只是注册路由；
// 评论：此处还省去了 Get、Post 等方法，包裹一层又何必呢？简洁
func (h *HTTPServer) AddRoute(method, path string, handler HandleFunc) {

}

// ServeHTTP HTTPServer 处理请求的入口
func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Req: r,
		Resp: w,
	}

	// 接下来，查找路由、执行业务逻辑
	h.serve(ctx)
}


func (h *HTTPServer) Start(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	
	/*
		这块儿做服务发现 - 注册本服务器到管理平台，譬如 etcd
	*/

	return http.Serve(ln, h)
}

func (h *HTTPServer) Start1(addr string) error {
	return http.ListenAndServe(addr, h)
	// 阻塞，如果拆分 listen 监听端口和 serve 服务器启动，这里没法做`服务发现`，对`框架使用者`而言。
}

func (h *HTTPServer) serve(ctx *Context) {

}



/* 步骤 1

	对于一个 web 框架而言，首先要有一个整体代表服务器的抽象，也就是 Server。

	Server 从特性上来说，至少要提供三部分功能：
	- 生命周期控制：即启动、关闭。如果在后期，我们还要考虑增加生命周期回调特性。
	- 路由注册接口：提供路由注册功能
	- 作为 http 包到 web 框架的桥梁
								
						
	req  ——>  |            
			  | http 包 | web 框架 | 业务
	resp <——  |            
					


	http 包暴露了一个接口，Handler
	它是我们引入自定义 web 框架相关的连接点
 
*/

