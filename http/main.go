/**
 * @Author: realpeanut
 * @Date: 2022/6/18 07:59
 */
package main

import (
	"net/http"
	"repeat/framework"
)

func main() {
	core := framework.NewCore()

	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":80",
	}
	server.ListenAndServe()
}
