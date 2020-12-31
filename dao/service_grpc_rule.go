package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	// "github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
)

type GrpcRule struct {
	Id 				int64 `json:"id" gorm:"primary_key"`
	ServiceId 		int64 `json:"service_id" gorm:"column:service_id" description:"服务id"`
	Port 			int `json:"port" gorm:"column:port" description:"端口"`
	HeaderTransfor 	string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换 支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (this *GrpcRule) TableName() string {
	return "gateway_service_grpc_rule"
}

// 从数据库查询数据
func (this *GrpcRule) Find(c *gin.Context, tx *gorm.DB, search *GrpcRule) (*GrpcRule, error) {
	out := &GrpcRule{}
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的日志
	err := tx.SetCtx(public.GetGinTraceContext(c)).
			  Where(search).
			  Find(out).
			  Error
	return out, err
}

// 保存数据到数据库
func (this *GrpcRule) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}