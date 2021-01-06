package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/contrib/sessions"
	"github.com/pkg/errors"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/yaolixiao/golang_common/lib"
	// "encoding/json"
	"fmt"
	"strings"
	"time"
)

type ServiceController struct {}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/service_delete", service.ServiceDelete)
	group.GET("/service_detail", service.ServiceDetail)
	group.GET("/service_stat", service.ServiceStat)
	group.POST("/service_add_http", service.ServiceAddHTTP)
	group.POST("/service_update_http", service.ServiceUpdateHTTP)

	group.POST("/service_add_grpc", service.ServiceAddGrpc)
	group.POST("/service_update_grpc", service.ServiceUpdateGrpc)
	group.POST("/service_add_tcp", service.ServiceAddTcp)
	group.POST("/service_update_tcp", service.ServiceUpdateTcp)
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

// ServiceDetail godoc
// @Summary 服务详情
// @Description 服务详情
// @Tags 服务管理接口
// @ID /service/service_detail
// @Accept  json
// @Produce  json
// @Param id query int64 true "服务id"
// @Success 200 {object} middleware.Response{data=dao.ServiceDetail} "success"
// @Router /service/service_detail [get]
func (this *ServiceController) ServiceDetail(c *gin.Context) {
	
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
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	
	middleware.ResponseSuccess(c, serviceDetail)
}

// ServiceStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 服务管理接口
// @ID /service/service_stat
// @Accept  json
// @Produce  json
// @Param id query int64 true "服务id"
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /service/service_stat [get]
func (this *ServiceController) ServiceStat(c *gin.Context) {
	
	// 校验参数
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// tx, err := lib.GetGormPool("dev")
	// if err != nil {
	// 	middleware.ResponseError(c, 2001, err)
	// 	return
	// }

	// // 读取基本信息
	// serviceInfo := &dao.ServiceInfo{Id: params.Id}
	// serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	// if err != nil {
	// 	middleware.ResponseError(c, 2002, err)
	// 	return
	// }

	// // 读取详情数据
	// serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo)
	// if err != nil {
	// 	middleware.ResponseError(c, 2003, err)
	// 	return
	// }

	// 开始进行服务统计
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

// ServiceAddHTTP godoc
// @Summary 添加HTTP服务
// @Description 添加HTTP服务
// @Tags 服务管理接口
// @ID /service/service_add_http
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_http [post]
func (this *ServiceController) ServiceAddHTTP(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceAddHTTPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 数据库对参数进行校验
	// 3.1 获取数据库连接池
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 3.2 验证服务名称是否存在
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	if _, err = serviceInfo.Find(c, tx, serviceInfo); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2003, errors.New("服务已存在"))
		return
	}

	// 3.3 验证域名前缀是否存在
	httpRule := &dao.HttpRule{RuleType: params.RuleType, Rule: params.Rule}
	if _, err = httpRule.Find(c, tx, httpRule); err == nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, errors.New("接入前缀或域名已存在"))
		return
	}

	// 4. 数据入库
	infoModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := infoModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}
	// 这里可以拿到服务id：infoModel.Id
	httpRuleModel := &dao.HttpRule{
		ServiceId: infoModel.Id,
		RuleType: params.RuleType,
		Rule: params.Rule,
		NeedHttps: params.NeedHttps,
		NeedStripUri: params.NeedStripUri,
		NeedWebsocket: params.NeedWebsocket,
		UrlRewrite: params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := httpRuleModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}
	AccessControlModel := &dao.AccessControl{
		ServiceId: infoModel.Id,
		OpenAuth: params.OpenAuth,
		BlackList: params.BlackList,
		WhiteList: params.WhiteList,
		// WhiteHostName: params.WhiteHostName,
		ClientipFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit: params.ServiceFlowLimit,
	}
	if err := AccessControlModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}
	LoadBalanceModel := &dao.LoadBalance{
		ServiceId: infoModel.Id,
		RoundType: params.RoundType,
		IpList: params.IpList,
		WeightList: params.WeightList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout: params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout: params.UpstreamIdleTimeout,
		UpstreamMaxIdle: params.UpstreamMaxIdle,
	}
	if err := LoadBalanceModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}
	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceUpdateHTTP godoc
// @Summary 修改HTTP服务
// @Description 修改HTTP服务
// @Tags 服务管理接口
// @ID /service/service_update_http
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_http [post]
func (this *ServiceController) ServiceUpdateHTTP(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceUpdateHTTPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 连接数据库，
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 4. 根据参数拿到服务详情（包含该服务的所有数据）
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, errors.New("服务不存在"))
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo);
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, errors.New("服务不存在"))
		return
	}

	// 5. 更新数据
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	httpRule := serviceDetail.HTTPRule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.NeedStripUri = params.NeedStripUri
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientipFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	loadBalance := serviceDetail.LoadBalance
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadBalance.UpstreamHeaderTimeout = params.UpstreamHeaderTimeout
	loadBalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadBalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}

	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceAddGrpc godoc
// @Summary 添加Grpc服务
// @Description 添加Grpc服务
// @Tags 服务管理接口
// @ID /service/service_add_grpc
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddGrpcInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_grpc [post]
func (this *ServiceController) ServiceAddGrpc(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceAddGrpcInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 数据库对参数进行校验
	// 3.1 获取数据库连接池
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 3.2 验证服务名称是否存在
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	if _, err = serviceInfo.Find(c, tx, serviceInfo); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2003, errors.New("服务已存在"))
		return
	}

	// 3.3 验证端口是否被占用（tcp）
	tcpRule := &dao.TcpRule{Port: params.Port}
	if _, err = tcpRule.Find(c, tx, tcpRule); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2004, errors.New("端口被占用"))
		return
	}

	// 3.4 验证端口是否被占用（grpc）
	grpcRule := &dao.GrpcRule{Port: params.Port}
	if _, err = grpcRule.Find(c, tx, grpcRule); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2005, errors.New("端口被占用"))
		return
	}

	// 4. 数据入库
	infoModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := infoModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}
	// 这里可以拿到服务id：infoModel.Id
	grpcRuleModel := &dao.GrpcRule{
		ServiceId: infoModel.Id,
		Port: params.Port,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := grpcRuleModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}
	AccessControlModel := &dao.AccessControl{
		ServiceId: infoModel.Id,
		OpenAuth: params.OpenAuth,
		BlackList: params.BlackList,
		WhiteList: params.WhiteList,
		WhiteHostName: params.WhiteHostName,
		ClientipFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit: params.ServiceFlowLimit,
	}
	if err := AccessControlModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}
	LoadBalanceModel := &dao.LoadBalance{
		ServiceId: infoModel.Id,
		RoundType: params.RoundType,
		IpList: params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := LoadBalanceModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2009, err)
		return
	}
	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceUpdateGrpc godoc
// @Summary 修改Grpc服务
// @Description 修改Grpc服务
// @Tags 服务管理接口
// @ID /service/service_update_grpc
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateGrpcInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_grpc [post]
func (this *ServiceController) ServiceUpdateGrpc(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceUpdateGrpcInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 连接数据库，
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 4. 根据参数拿到服务详情（包含该服务的所有数据）
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, errors.New("服务不存在"))
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo);
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, errors.New("服务不存在"))
		return
	}

	// 5. 更新数据
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	grpcRule := serviceDetail.GRPCRule
	grpcRule.ServiceId = info.Id
	grpcRule.Port = params.Port
	grpcRule.HeaderTransfor = params.HeaderTransfor
	if err := grpcRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientipFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	loadBalance := serviceDetail.LoadBalance
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}

	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceAddTcp godoc
// @Summary 添加Tcp服务
// @Description 添加Tcp服务
// @Tags 服务管理接口
// @ID /service/service_add_tcp
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddTcpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_tcp [post]
func (this *ServiceController) ServiceAddTcp(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceAddTcpInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 数据库对参数进行校验
	// 3.1 获取数据库连接池
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 3.2 验证服务名称是否存在
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	if _, err = serviceInfo.Find(c, tx, serviceInfo); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2003, errors.New("服务已存在"))
		return
	}

	// 3.3 验证端口是否被占用（tcp）
	tcpRule := &dao.TcpRule{Port: params.Port}
	if _, err = tcpRule.Find(c, tx, tcpRule); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2004, errors.New("端口被占用"))
		return
	}

	// 3.4 验证端口是否被占用（grpc）
	grpcRule := &dao.GrpcRule{Port: params.Port}
	if _, err = grpcRule.Find(c, tx, grpcRule); err == nil {
		tx.Rollback()
		// github.com/pkg/errors 可以携带错误堆栈
		middleware.ResponseError(c, 2005, errors.New("端口被占用"))
		return
	}

	// 4. 数据入库
	infoModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := infoModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}
	// 这里可以拿到服务id：infoModel.Id
	tcpRuleModel := &dao.TcpRule{
		ServiceId: infoModel.Id,
		Port: params.Port,
	}
	if err := tcpRuleModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}
	AccessControlModel := &dao.AccessControl{
		ServiceId: infoModel.Id,
		OpenAuth: params.OpenAuth,
		BlackList: params.BlackList,
		WhiteList: params.WhiteList,
		WhiteHostName: params.WhiteHostName,
		ClientipFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit: params.ServiceFlowLimit,
	}
	if err := AccessControlModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}
	LoadBalanceModel := &dao.LoadBalance{
		ServiceId: infoModel.Id,
		RoundType: params.RoundType,
		IpList: params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := LoadBalanceModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2009, err)
		return
	}
	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceUpdateTcp godoc
// @Summary 修改Tcp服务
// @Description 修改Tcp服务
// @Tags 服务管理接口
// @ID /service/service_update_tcp
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateTcpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_tcp [post]
func (this *ServiceController) ServiceUpdateTcp(c *gin.Context) {

	// 1. 中间件对参数进行校验，并绑定数据
	params := &dto.ServiceUpdateTcpInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 验证IP列表和权重列表数量是否相等
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表数量不一致"))
		return
	}

	// 3. 连接数据库，
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin() //开启事务，和下面的tx.Rollback()一起使用

	// 4. 根据参数拿到服务详情（包含该服务的所有数据）
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, errors.New("服务不存在"))
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo);
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, errors.New("服务不存在"))
		return
	}

	// 5. 更新数据
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	tcpRule := serviceDetail.TCPRule
	tcpRule.ServiceId = info.Id
	tcpRule.Port = params.Port
	if err := tcpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientipFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	loadBalance := serviceDetail.LoadBalance
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}

	// 提交保存数据库
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}