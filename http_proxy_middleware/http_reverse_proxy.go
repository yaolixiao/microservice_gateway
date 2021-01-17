package http_proxy_middleware

import (
	"github.com/yaolixiao/microservice_gateway/reverse_proxy"
	"github.com/gin-gonic/gin"
)

// 根据请求信息，匹配接入方式
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		reverse_proxy.NewMultipleHostsReverseProxy(lb loadbalance.LoadBalance)
	}
}
