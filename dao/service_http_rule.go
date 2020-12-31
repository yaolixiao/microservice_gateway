package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	// "github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
)

type HttpRule struct {
	Id 				int64 `json:"id" gorm:"primary_key"`
	ServiceId 		int64 `json:"service_id" gorm:"column:service_id" description:"服务id"`
	RuleType 		int `json:"rule_type" gorm:"column:rule_type" description:"匹配类型 domain=域名, url_prefix=url前缀"`
	Rule 			string `json:"rule" gorm:"column:rule" description:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps 		int `json:"need_https" gorm:"column:need_https" description:"支持https 1=支持"`
	NeedStripUri 	int `json:"need_strip_uri" gorm:"column:need_strip_uri" description:"启用strip_uri 1=启用"`
	NeedWebsocket 	int `json:"need_websocket" gorm:"column:need_websocket" description:"启用websocket 1=启用"`
	UrlRewrite 		string `json:"url_rewrite" gorm:"column:url_rewrite" description:"url重写功能，每行一个"`
	HeaderTransfor 	string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换 支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (this *HttpRule) TableName() string {
	return "gateway_service_http_rule"
}

// 从数据库查询数据
func (this *HttpRule) Find(c *gin.Context, tx *gorm.DB, search *HttpRule) (*HttpRule, error) {
	out := &HttpRule{}
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的日志
	err := tx.SetCtx(public.GetGinTraceContext(c)).
			  Where(search).
			  Find(out).
			  Error
	return out, err
}

// 保存数据到数据库
func (this *HttpRule) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}