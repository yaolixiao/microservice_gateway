package dto

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/public"
)

type AdminInfoOutput struct {
	Id 				int 		`json:"id"`
	Name 			string 		`json:"name"`
	LoginTime 		time.Time 	`json:"login_time"`
	Avatar 			string 		`json:"avatar"`
	Introduction 	string 		`json:"introduction"`
	Roles 			[]string 	`json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

// 校验参数
func (this *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, this)
}