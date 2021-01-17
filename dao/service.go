package dao

import (
	"errors"
	"net/http/httptest"
	"strings"
	"sync"

	"github.com/yaolixiao/microservice_gateway/public"

	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/golang_common/lib"
	"github.com/yaolixiao/microservice_gateway/dto"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http_rule" description:""`
	TCPRule       *TcpRule       `json:"tcp_rule" description:""`
	GRPCRule      *GrpcRule      `json:"grpc_rule" description:""`
	LoadBalance   *LoadBalance   `json:"load_balance" description:""`
	AccessControl *AccessControl `json:"access_control" description:""`
}

var ServiceManagerHandler *ServiceManager

func init() {
	ServiceManagerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap   map[string]*ServiceDetail
	ServiceSlice []*ServiceDetail
	Locker       sync.RWMutex
	init         sync.Once
	err          error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:   map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker:       sync.RWMutex{},
		init:         sync.Once{},
	}
}

// 接入匹配的方法
func (this *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail, error) {

	// 前缀匹配 /abc == this.ServiceSlice.rule
	// 域名匹配 www.test.com == this.ServiceSlice.rule
	// host c.Request.Host
	// path c.Request.URL.Path

	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	path := c.Request.URL.Path

	for _, serviceitem := range this.ServiceSlice {
		if serviceitem.Info.LoadType != public.LoadTypeHTTP {
			continue
		}
		if serviceitem.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			if serviceitem.HTTPRule.Rule == host {
				return serviceitem, nil
			}
		}
		if serviceitem.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL {
			if strings.HasPrefix(path, serviceitem.HTTPRule.Rule) {
				return serviceitem, nil
			}
		}
	}

	return nil, errors.New("not matched service.")
}

// 只执行一次，将服务加载到内存
func (this *ServiceManager) LoadOnce() error {
	this.init.Do(func() {

		tx, err := lib.GetGormPool("dev")
		if err != nil {
			this.err = err
			return
		}

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		serviceInfo := &ServiceInfo{}
		params := &dto.ServiceListInput{PageNo: 1, PageSize: 9999}
		list, _, err := serviceInfo.PageList(c, tx, params)
		if err != nil {
			this.err = err
			return
		}

		this.Locker.Lock()
		defer this.Locker.Unlock()
		for _, listItem := range list {
			tempitem := listItem
			serviceDetail, err := tempitem.ServiceDetail(c, tx, &tempitem)
			if err != nil {
				this.err = err
				return
			}

			this.ServiceMap[tempitem.ServiceName] = serviceDetail
			this.ServiceSlice = append(this.ServiceSlice, serviceDetail)
		}
	})
	return this.err
}
