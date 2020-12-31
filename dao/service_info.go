package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
	"time"
	// "errors"
)

type ServiceInfo struct {
	Id int64 `json:"id" gorm:"primary_key" description:"自增主键"`
	ServiceName string `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	LoadType int `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"是否已删除 0:否 1:是"`
}

func (this *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

// 服务列表分页查询
func (this *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	// 定义并初始化返回值
	total := int64(0)
	list := []ServiceInfo{}
	// 查询页数偏移量
	offset := (param.PageNo - 1) * param.PageSize

	query := tx.SetCtx(public.GetGinTraceContext(c))

	// 设置 Table(this.TableName()) 是为了下面查询总数 Count() 的需要
	// is_delete=0 保证查出的都是有效数据
	query = query.Table(this.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)", "%" + param.Info + "%", "%" + param.Info + "%")
	}
	
	err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; 
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	// 查询总数
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

// 服务详情查询
func (this *ServiceInfo) ServiceDetail(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	httpRule := &HttpRule{ServiceId: search.Id}
	httpRule, err := httpRule.Find(c, tx, httpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	tcpRule := &TcpRule{ServiceId: search.Id}
	tcpRule, err = tcpRule.Find(c, tx, tcpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	grpcRule := &GrpcRule{ServiceId: search.Id}
	grpcRule, err = grpcRule.Find(c, tx, grpcRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	loadBalance := &LoadBalance{ServiceId: search.Id}
	loadBalance, err = loadBalance.Find(c, tx, loadBalance)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	accessControl := &AccessControl{ServiceId: search.Id}
	accessControl, err = accessControl.Find(c, tx, accessControl)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	detail := &ServiceDetail{
		Info: search,
		HTTPRule: httpRule,
		TCPRule: tcpRule,
		GRPCRule: grpcRule,
		LoadBalance: loadBalance,
		AccessControl: accessControl,
	}
	return detail, nil
}

// 从数据库查询数据
func (this *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
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
func (this *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}