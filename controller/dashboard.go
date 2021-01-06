package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/yaolixiao/golang_common/lib"
	"time"
)

type DashboardController struct {}

func DashboardRegister(group *gin.RouterGroup) {
	dash := &DashboardController{}
	group.GET("/panel_group_data", dash.PanelGroupData)
	group.GET("/flow_stat", dash.FlowStat)
	group.GET("/service_stat", dash.ServiceStat)
}

// PanelGroupData godoc
// @Summary 大盘指标统计
// @Description 大盘指标统计
// @Tags 首页大盘
// @ID /dashboard/panel_group_data
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput} "success"
// @Router /dashboard/panel_group_data [get]
func (this *DashboardController) PanelGroupData(c *gin.Context) {
	
	// 取出 ServiceNum
	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(c, lib.GORMDefaultPool, &dto.ServiceListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// 取出 AppNum
	appInfo := &dao.App{}
	_, appNum, err := appInfo.AppList(c, lib.GORMDefaultPool, &dto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	out := &dto.PanelGroupDataOutput{
		ServiceNum: serviceNum,
		AppNum: appNum,
		CurrentQps: 0,
		TodayRequestNum: 0,
	}
	middleware.ResponseSuccess(c, out)
}

// FlowStat godoc
// @Summary 今日昨日统计
// @Description 今日昨日统计
// @Tags 首页大盘
// @ID /dashboard/flow_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /dashboard/flow_stat [get]
func (this *DashboardController) FlowStat(c *gin.Context) {

	// 开始进行统计
	todayList := []int64{}
	for i := 0; i <= time.Now().Hour(); i++ {
		todayList = append(todayList, 0)
	}
	yesterdayList := []int64{}
	for i := 0; i <= 23; i++ {
		yesterdayList = append(yesterdayList, 0)
	}
	
	middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
		Today: todayList,
		Yesterday: yesterdayList,
	})
}

// ServiceStat godoc
// @Summary 服务类型统计
// @Description 服务类型统计
// @Tags 首页大盘
// @ID /dashboard/service_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.DashServiceStatOutput} "success"
// @Router /dashboard/service_stat [get]
func (this *DashboardController) ServiceStat(c *gin.Context) {
	
	serviceInfo := &dao.ServiceInfo{}
	list, err := serviceInfo.GroupByLoadType(c, lib.GORMDefaultPool)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	legend := []string{}
	for index, item := range list {
		name, ok := public.LoadTypeMap[item.LoadType]
		if !ok {
			middleware.ResponseError(c, 2002, errors.New("load_type not found."))
			return
		}
		list[index].Name = name
		legend = append(legend, name)
	}

	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data: list,
	}
	middleware.ResponseSuccess(c, out)
}