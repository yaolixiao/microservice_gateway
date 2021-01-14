package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yaolixiao/microservice_gateway/basic/proxy"
	"github.com/yaolixiao/microservice_gateway/basic/proxy/loadbalance"
	"github.com/yaolixiao/microservice_gateway/basic/proxy/middleware"
)

// 使用数组构建中间件-思路
// 组装：构建中间件URL路由 -> 构建URL的中间件方法数组 -> 使用use方法整合路由与方法数组
// 调用：构建方法请求逻辑 -> 封装 http.Handler接口与http.Server整合

var addr = "127.0.0.1:2001"

func main() {

	// 创建反向代理服务器 (默认端口2002)
	lb := loadbalance.LoadBalanceFactory(loadbalance.RANDOM)
	lb.Add("http://127.0.0.1:2003")
	lb.Add("http://127.0.0.1:2004")
	reverseProxy := proxy.NewMultipleHostsReverseProxy(lb)

	// 1. 【组装】
	// 初始化方法数组路由器
	sliceRouter := middleware.NewSliceRouter()

	// 添加路由，给路由绑定中间件
	sliceRouter.Group("/base").Use(middleware.TraceLogSlice(), func(c *middleware.SliceRouterContext) {
		// 中间件处理业务逻辑
		c.Rw.Write([]byte("业务逻辑..."))
	})
	sliceRouter.Group("/").Use(middleware.TraceLogSlice(), func(c *middleware.SliceRouterContext) {
		// 中间件请求到代理服务器
		fmt.Println("reverseProxy...")
		reverseProxy.ServeHTTP(c.Rw, c.Req)
	})

	// 2. 【调用】
	// 构建Handler方法
	routerHandler := middleware.NewSliceRouterHandler(nil, sliceRouter)
	log.Println("starting server at", addr)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
