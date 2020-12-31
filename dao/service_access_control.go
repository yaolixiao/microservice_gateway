package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	// "github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
)

type AccessControl struct {
	Id 					int64 `json:"id" gorm:"primary_key"`
	ServiceId 			int64 `json:"service_id" gorm:"column:service_id" description:"服务id"`
	OpenAuth 			int `json:"open_auth" gorm:"column:open_auth" description:"是否开启权限 1=开启"`
	BlackList 			string `json:"black_list" gorm:"column:black_list" description:"黑名单ip"`
	WhiteList 			string `json:"white_list" gorm:"column:white_list" description:"白名单ip"`
	WhiteHostName 		string `json:"white_host_name" gorm:"column:white_host_name" description:"白名单主机"`
	ClientipFlowLimit 	int `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流"`
	ServiceFlowLimit 	int `json:"service_flow_limit" gorm:"column:service_flow_limit" description:"服务限流"`
}

func (this *AccessControl) TableName() string {
	return "gateway_service_access_control"
}

// 从数据库查询数据
func (this *AccessControl) Find(c *gin.Context, tx *gorm.DB, search *AccessControl) (*AccessControl, error) {
	out := &AccessControl{}
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的日志
	err := tx.SetCtx(public.GetGinTraceContext(c)).
			  Where(search).
			  Find(out).
			  Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// 保存数据到数据库
func (this *AccessControl) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}