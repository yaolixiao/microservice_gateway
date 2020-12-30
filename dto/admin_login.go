package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	Id 			int 		`json:"id"`
	UserName 	string 		`json:"username"`
	LoginTime 	time.Time 	`json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,is-validuser-username"`//管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`//密码
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`//token
}

// 绑定结构体，校验参数
func (this *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, this)
}