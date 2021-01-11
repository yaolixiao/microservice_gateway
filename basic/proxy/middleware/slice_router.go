package middleware

import (
	"context"
	"math"
	"net/http"
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
	path string
	handers []HandlerFunc
}

type SliceRouterContext struct {
	Rw http.ResponseWriter
	Req http.Request
	Ctx context.Context
	*SliceGroup
	index int8
}

// func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouterContext {
// 	newSliceGroup := &SliceGroup{}

// 	// 最长url前缀匹配
// 	matchUrlLen := 0
// 	for _, group := range r.groups {
// 		if strings.HasPrefix(req.RequestURI, group.path) {
// 			pathLen := len(group.path)
// 			if pathLen > matchUrlLen {
// 				matchUrlLen = pathLen
// 				*newSliceGroup = *group
// 			}
// 		}
// 	}

// 	c := &SliceRouterContext{Rw: rw, Req: req, Ctx: req.Context(), SliceGroup: newSliceGroup}
// 	c.Reset()
// 	return c
// }

// func (this *SliceRouterContext) Get(key interface{}) interface{} {
// 	return this.Ctx.Value(key)
// }

// func (this *SliceRouterContext) Set(key, val interface{}) {
// 	this.Ctx = context.WithValue(this.Ctx, key, val)
// }
