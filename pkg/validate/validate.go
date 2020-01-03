package validate

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var trans ut.Translator

// SetUp 注册汉化 与自定义验证规则
func SetUp() {
	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Validate错误提示汉化
		zhCn := zh.New()
		uni := ut.New(zhCn, zhCn)
		trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(v, trans)

		// 注册验证规则
		v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
			regular := "^1[3456789]\\d{9}$"
			reg := regexp.MustCompile(regular)
			if reg.MatchString(fl.Field().String()) {
				return true
			}
			return false
		})

		// 注册错误信息
		v.RegisterTranslation("password", trans, func(ut ut.Translator) error {
			return ut.Add("password", "{0}不符合规则", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		})

		// 注册验证规则
		v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			regular := "^.{6,18}$"
			reg := regexp.MustCompile(regular)
			if reg.MatchString(fl.Field().String()) {
				return true
			}
			return false
		})

		// 注册错误信息
		v.RegisterTranslation("password", trans, func(ut ut.Translator) error {
			return ut.Add("password", "{0}不符合规则", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		})
	}
}

// ValidatorError 错误文本转换
func ValidatorError(err error) string {
	var str string
	m := FieldTrans()

	errs, ok := err.(validator.ValidationErrors)

	// 如果是json绑定错误 直接返回
	if !ok {
		return err.Error()
	}

	for _, e := range errs {
		transtr := e.Translate(trans)
		// 判断错误字段是否在命名中，如果在，则替换错误信息中的字段
		if rp, ok := m[e.Field()]; ok {
			str = strings.Replace(transtr, e.Field(), rp, 1)
		} else {
			str = transtr
		}
	}

	// 返回错误信息
	return str
}

// ModelFieldTran 模型名称转换
type ModelFieldTran map[string]string

// FieldTrans 模型字段转换
func FieldTrans() ModelFieldTran {
	m := ModelFieldTran{}
	m["Mobile"] = "手机号"
	m["Password"] = "密码"
	m["Email"] = "邮箱"
	return m
}

// 常用正则
// 用户名：  	/^[a-z0-9_-]{3,16}$/
// 密码：	    /^[a-z0-9_-]{6,18}$/
// 十六进制值：	/^#?([a-f0-9]{6}|[a-f0-9]{3})$/
// 电子邮箱	：  /^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$/
// 		    /^[a-z\d]+(\.[a-z\d]+)*@([\da-z](-[\da-z])?)+(\.{1,2}[a-z]+)+$/
// URL： 	    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
// IP 地址：	/((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)/
//             /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
// HTML 标签：	/^<([a-z]+)([^<]+)*(?:>(.*)<\/\1>|\s+\/>)$/

// 删除代码\\注释：      	(?<!http:|\S)//.*$
// Unicode编码中的汉字范围：	/^[\u2E80-\u9FFF]+$/
