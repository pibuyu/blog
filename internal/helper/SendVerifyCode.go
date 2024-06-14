package helper

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendVerifyCode 邮箱发送验证码
func SendVerifyCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "<3531095171@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "blog_api_rpc系统验证码"
	e.HTML = []byte("<b>" + "您的验证码是：" + code + "</b>")
	//返回EOF的时候，关闭SSL重试
	return e.SendWithTLS(
		"smtp.qq.com:465",
		smtp.PlainAuth("",
			"3531095171@qq.com",
			"hyevsfluzhfldbhc",
			"smtp.qq.com"),
		&tls.Config{
			ServerName:         "smtp.qq.com",
			InsecureSkipVerify: true,
		},
	)
}
