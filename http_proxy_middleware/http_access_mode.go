package http_proxy_middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
)

// 根据请求信息，匹配接入方式
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		fmt.Println("matched service:", public.Obj2Json(service))
		// 设置服务到上下文，方便下游中间件取到服务信息
		c.Set("service", service)
		c.Next()
	}
}
