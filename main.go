package main

import (
	"fmt"
	"net/http"

	"github.com/sjxiang/one/web"
)

func main() {
	w := web.HTTPServer{}
	w.AddRoute(http.MethodGet, "/ping", func(ctx *web.Context) {
		// ctx.Resp.Write([]byte("pong"))
		fmt.Println("pong")
	})

	w.Start(":8080")
}

// 前 25 min，讲面试