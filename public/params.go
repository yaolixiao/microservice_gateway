package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	// c.ShouldBind:将参数数据绑定到params
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	// 获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	// 获取翻译器
	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}
	// 验证是否符合结构体tag设置的规则
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ", "))
	}
	return nil
}

// 获取验证器（验证器通过中间件translation初始化）
func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(ValidatorKey);
	if !ok {
		return nil, errors.New("未设置验证器")
	}
	validator, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validator, nil
}

// 获取翻译器（翻译器通过中间件translation初始化）
func GetTranslation(c *gin.Context) (ut.Translator, error) {
	val, ok := c.Get(TranslatorKey);
	if !ok {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := val.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}