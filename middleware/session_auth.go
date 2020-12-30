package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/public"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		info, ok := session.Get(public.AdminSessionInfoKey).(string)
		if !ok || info == "" {
			ResponseError(c, InternalErrorCode, errors.New("user not login."))
			c.Abort()
			return
		}
		c.Next()
	}
}