package validate

import (
	"errors"
	"regexp"
)

// VerifyMobile 验证电话号码
func VerifyMobile(mobile string) error {
	regular := "^1[3456789]\\d{9}$"
	reg := regexp.MustCompile(regular)
	if reg.MatchString(mobile) {
		return nil
	}
	return errors.New("手机号码不符合要求")
}

// VerifyPwd 验证密码
func VerifyPwd(password string) error {
	regular := "^[a-z0-9_-]{6,18}$"
	reg := regexp.MustCompile(regular)
	if reg.MatchString(password) {
		return nil
	}
	return errors.New("密码不符合要求")
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
