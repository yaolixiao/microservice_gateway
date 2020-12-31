package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	// "github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
)

type TcpRule struct {
	Id 			int64 `json:"id" gorm:"primary_key"`
	ServiceId 	int64 `json:"service_id" gorm:"column:service_id" description:"服务id"`
	Port 		int `json:"port" gorm:"column:port" description:"端口"`
}

func (this *TcpRule) TableName() string {
	return "gateway_service_tcp_rule"
}

// 从数据库查询数据
func (this *TcpRule) Find(c *gin.Context, tx *gorm.DB, search *TcpRule) (*TcpRule, error) {
	out := &TcpRule{}
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的日志
	err := tx.SetCtx(public.GetGinTraceContext(c)).
			  Where(search).
			  Find(out).
			  Error
	return out, err
}

// 保存数据到数据库
func (this *TcpRule) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}