package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/microservice_gateway/public"
)

type ServiceListInput struct {
	Info string `json:"info" form:"info" comment:"关键词" example:"" validate:""`
	PageNo int `json:"page_no" form:"page_no" comment:"页面" example:"1" validate:"required"`
	PageSize int `json:"page_size" form:"page_size" comment:"每页多少条" example:"20" validate:"required"`
}

// 校验参数
func (this *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, this)
}

type ServiceDeleteInput struct {
	Id int64 `json:"id" form:"id" comment:"服务id" example:"" validate:"required"`
}

// 校验参数
func (this *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, this)
}

type ServiceListOutput struct {
	Total int64 `json:"total" form:"total" comment:"总数"`
	List []ServiceListItemOutput `json:"list" form:"list" comment:"服务列表"`
}

type ServiceListItemOutput struct {
	Id int64 `json:"id" form:"id"`//id
	ServiceName string `json:"service_name" form:"service_name"`//服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc"`//服务描述
	LoadType int `json:"load_type" form:"load_type"`//类型
	ServiceAddr string `json:"service_addr" form:"service_addr"`//服务地址
	Qps int64 `json:"qps" form:"qps"`//qps
	Qpd int64 `json:"qpd" form:"qpd"`//qpd
	TotalNode int `json:"total_node" form:"total_node"`//节点数
}