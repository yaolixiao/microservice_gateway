package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/http_proxy_middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping success",
		})
	})

	router.Use(http_proxy_middleware.HTTPAccessModeMiddleware())
	return router
}
