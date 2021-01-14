package middleware

import (
	"context"
	"math"
	"net/http"
	"strings"
	// "strings"
)

//目标定位是 tcp、http通用的中间件

const AbortIndex int8 = math.MaxInt8 / 2 // 最多63个中间件

type HandlerFunc func(*SliceRouterContext)

type SliceRouter struct {
	groups []*SliceGroup
}

type SliceGroup struct {
	*SliceRouter
	path    string
	handers []HandlerFunc
}

type SliceRouterContext struct {
	Rw  http.ResponseWriter
	Req *http.Request
	Ctx context.Context
	*SliceGroup
	index int8 // 控制中间件的执行顺序
}

// 1.1 创建路由器
func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}

// 1.2 给路由器添加group方法，以便于注册路由
func (this *SliceRouter) Group(path string) *SliceGroup {
	return &SliceGroup{
		SliceRouter: this,
		path:        path,
	}
}

// 1.3 给group添加use方法，用来添加中间件数组
func (this *SliceGroup) Use(middlewares ...HandlerFunc) *SliceGroup {
	// middlewares 是一组回调函数，函数参数需要router的上下文
	// router上下文是在调用过程中，通过handler的ServeHTTP方法来创建的

	// 绑定中间件
	this.handers = append(this.handers, middlewares...)

	// 将group注册到父级路由
	// 注册前先判断父级路由是否已存在该group
	var existsFlag = false
	for _, oldgroup := range this.SliceRouter.groups {
		if oldgroup == this {
			existsFlag = true
		}
	}
	if !existsFlag {
		this.SliceRouter.groups = append(this.SliceRouter.groups, this)
	}

	return this
}

// 2.1 构建Handler方法
// 创建结构体
type SliceRouterHandler struct {
	coreFunc func(*SliceRouterContext) http.Handler
	router   *SliceRouter
}

// 初始化Handler
func NewSliceRouterHandler(coreFunc func(*SliceRouterContext) http.Handler,
	router *SliceRouter) *SliceRouterHandler {
	return &SliceRouterHandler{
		coreFunc: coreFunc,
		router:   router,
	}
}

// 2.2 实现http.Handler接口，也就是实现ServeHTTP方法
func (this *SliceRouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 创建router上下文
	c := newSliceRouterContext(w, r, this.router)

	// 首先执行coreFunc, 默认中间件
	if this.coreFunc != nil {
		c.handers = append(c.handers, func(c *SliceRouterContext) {
			this.coreFunc(c).ServeHTTP(w, r)
		})
	}
	c.Reset()
	c.Next()
}

// 创建router上下文
func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouterContext {

	// 定义本次请求匹配到的路由
	newSliceGroup := &SliceGroup{}

	// 最长url前缀匹配
	matchUrlLen := 0
	for _, group := range r.groups {
		if strings.HasPrefix(req.RequestURI, group.path) {
			pathLen := len(group.path)
			if pathLen > matchUrlLen {
				matchUrlLen = pathLen
				*newSliceGroup = *group
			}
		}
	}

	c := &SliceRouterContext{Rw: rw, Req: req, Ctx: req.Context(), SliceGroup: newSliceGroup}
	c.Reset()
	return c
}

func (this *SliceRouterContext) Get(key interface{}) interface{} {
	return this.Ctx.Value(key)
}

func (this *SliceRouterContext) Set(key, val interface{}) {
	this.Ctx = context.WithValue(this.Ctx, key, val)
}

func (this *SliceRouterContext) Next() {
	this.index++
	for this.index < int8(len(this.handers)) {
		this.handers[this.index](this)
		this.index++
	}
}

// 跳出中间件方法
func (this *SliceRouterContext) Abort() {
	this.index = AbortIndex
}

// 是否跳过了回调
func (this *SliceRouterContext) IsAborted() bool {
	return this.index >= AbortIndex
}

func (this *SliceRouterContext) Reset() {
	this.index = -1
}
