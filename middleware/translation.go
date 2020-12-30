package middleware

import (
	"github.com/yaolixiao/microservice_gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

// translation中间件
// 功能：设置参数验证器、设置多语言翻译器
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		// 设置支持语言
		en := en.New()
		zh := zh.New()

		// 设置国际化翻译器
		uni := ut.New(zh, zh, en) // 第一个参数是默认支持语言

		// 设置参数验证器
		val := validator.New()

		// 根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		// 将翻译器注册到validator
		switch locale {
			case "en":
				en_translations.RegisterDefaultTranslations(val, trans)
				val.RegisterTagNameFunc(func(fld reflect.StructField) string {
					return fld.Tag.Get("en_comment")
				})
				break
			default:
				zh_translations.RegisterDefaultTranslations(val, trans)
				val.RegisterTagNameFunc(func(fld reflect.StructField) string {
					return fld.Tag.Get("comment")
				})

				//自定义验证方法
				//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
				val.RegisterValidation("is-validuser-username", func(fl validator.FieldLevel) bool {
					return fl.Field().String() == "admin"
				})

				//自定义翻译方法
				//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
				val.RegisterTranslation("is-validuser-username", trans, func(ut ut.Translator) error {
					return ut.Add("is-validuser-username", "{0} 填写不正确哦", true)
				}, func(ut ut.Translator, fe validator.FieldError) string {
					t, _ := ut.T("is-validuser-username", fe.Field())
					return t
				})
				break
		}
		c.Set(public.ValidatorKey, val)
		c.Set(public.TranslatorKey, trans)
		c.Next()
	}
}