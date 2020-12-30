package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/yaolixiao/golang_common/lib"
	"encoding/json"
	"fmt"
)

type AdminController struct {}

func AdminRegister(group *gin.RouterGroup) {
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
	// 修改密码（因为需要做登录验证，所以要在使用了session_auth中间件的admin这个路由下）
	group.POST("/change_pwd", admin.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (this *AdminController) AdminInfo(c *gin.Context) {

	// 1. 取出session信息，反序列化，为返回客户端准备数据
	session := sessions.Default(c)
	info := session.Get(public.AdminSessionInfoKey)
	sessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(info)), sessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 按照客户端要求组装数据
	out := &dto.AdminInfoOutput{
		Id: sessionInfo.Id,
		Name: sessionInfo.UserName,
		LoginTime: sessionInfo.LoginTime,
		Avatar: "http://.../avator.png",
		Introduction: "session info...",
		Roles: []string{"admin"},
	}

	// 3. 返回客户端
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (this *AdminController) ChangePwd(c *gin.Context) {

	// 1. 校验参数
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 取出session信息，反序列化
	session := sessions.Default(c)
	info := session.Get(public.AdminSessionInfoKey)
	sessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(info)), sessionInfo); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// 3. 根据session信息，查询数据库，得到adminInfo
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{UserName: sessionInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	// 4. 生成新密码：adminInfo.salt + params.Password sha256
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)

	// 5. 保存数据库
	// 说明：这里不需要设置更新时间
	// 原因：更新时间和创建时间 如果字段名是 UpdateAt和CreateAt，那么gorm会自动保存时间到数据库
	adminInfo.Password = saltPassword
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}

	// 返回客户端
	middleware.ResponseSuccess(c, "")
}