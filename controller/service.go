package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/contrib/sessions"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/yaolixiao/golang_common/lib"
	// "encoding/json"
	"fmt"
)

type ServiceController struct {}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/service_delete", service.ServiceDelete)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理接口
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_no query int true "页码"
// @Param page_size query int true "每页多少条"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (this *ServiceController) ServiceList(c *gin.Context) {
	
	// 校验参数
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	outList := []dto.ServiceListItemOutput{}
	for _, listItem := range list {

		serviceDetail, err := listItem.ServiceDetail(c, tx, &listItem)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}

		// 定义ServiceAddr，有3种接入方式
		// 1. http后缀接入：clusterIP + clusterPort + path
		// 2. http域名接入：domain
		// 3. tcp、grpc接入：clusterIP + servicePort
		serviceAddr := "unknow"
		clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && 
		   serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL {
			   if serviceDetail.HTTPRule.NeedHttps == 1 {
				   serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterSSLPort, serviceDetail.HTTPRule.Rule)
			   }
			   if serviceDetail.HTTPRule.NeedHttps == 0 {
				   serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterPort, serviceDetail.HTTPRule.Rule)
			   }
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && 
		   serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			   serviceAddr = serviceDetail.HTTPRule.Rule
		}
		if serviceDetail.Info.LoadType == public.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TCPRule.Port)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
		}
		ipList := serviceDetail.LoadBalance.GetIPListByModel()
		outItem := dto.ServiceListItemOutput{
			Id: listItem.Id,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			LoadType: listItem.LoadType,
			ServiceAddr: serviceAddr,
			Qps: 0,
			Qpd: 0,
			TotalNode: len(ipList),
		}
		outList = append(outList, outItem)
	}

	out := &dto.ServiceListOutput{
		Total: total,
		List: outList,
	}
	middleware.ResponseSuccess(c, out)
}

// ServiceDelete godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理接口
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query int64 true "服务id"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_delete [get]
func (this *ServiceController) ServiceDelete(c *gin.Context) {
	
	// 校验参数
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{Id: params.Id}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	serviceInfo.IsDelete = 1
	if err := serviceInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	
	middleware.ResponseSuccess(c, "")
}