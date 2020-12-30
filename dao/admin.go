package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/yaolixiao/gorm"
	"github.com/yaolixiao/microservice_gateway/dto"
	"github.com/yaolixiao/microservice_gateway/public"
	"time"
	"errors"

)

type Admin struct {
	Id int `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName string `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Password string `json:"password" gorm:"column:password" description:"密码"`
	Salt string `json:"salt" gorm:"column:salt" description:"盐"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (this *Admin) TableName() string {
	return "gateway_admin"
}

// 密码加密、验证
// 1. 根据 params.UserName 从数据库中取得管理员信息 admininfo
// 2. 密码加密，admininfo.salt + params.Password 一起sha256加密 => saltPassword
// 3. 验证密码是否相等 saltPassword == admininfo.Password ?
func (this *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	admininfo, err := this.Find(c, tx, &Admin{UserName: param.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}

	saltPassword := public.GenSaltPassword(admininfo.Salt, param.Password)
	if admininfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}

	return admininfo, nil
}

// 从数据库中取得管理员信息
func (this *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
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

// 保存管理员信息
func (this *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	// tx.SetCtx设置ginTrace上下文的好处是：可以在log日志打印mysql的链路日志
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(this).Error
}