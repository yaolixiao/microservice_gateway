package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/yaolixiao/microservice_gateway/dao"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/middleware"
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/yaolixiao/golang_common/lib"
	"time"
	"encoding/json"
)

type AdminLoginController struct {}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (this *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 获取【mysql】数据库连接池
	tx, err := lib.GetGormPool("dev")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// 登录验证
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	// 登录验证通过，保存session
	sessionInfo := &dto.AdminSessionInfo{
		Id: 		admin.Id,
		UserName: 	admin.UserName,
		LoginTime: 	time.Now(),
	}
	data, err := json.Marshal(sessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	session := sessions.Default(c)
	session.Set(public.AdminSessionInfoKey, string(data))
	session.Save()

	// 输出
	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLoginOut godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (this *AdminLoginController) AdminLoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(public.AdminSessionInfoKey)
	session.Save()
	middleware.ResponseSuccess(c, "")
}