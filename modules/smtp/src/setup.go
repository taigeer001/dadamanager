package smtp

import (
	"fmt"
	"strings"
	"net/smtp"
)

var auth = NewMyAuth()

type MyAuth struct {
	host string
	port string
	username string
	passwd string
}

func NewMyAuth() MyAuth {
	return MyAuth{
		"smtp.qq.com",
		"25",
		"904037433@qq.com",
		"zr11235813",
	}
}
func (this *MyAuth) PlainAuth() smtp.Auth {
	return smtp.PlainAuth(
		"",
		this.username,
		this.passwd,
		this.host,
	)
}

func SetAuth()  {
	
}

func SendMail(addr string, sub string, content string) error {
	sub = fmt.Sprintf("subject: %s\r\n\r\n", sub)
	mailList := strings.Split( addr,",")
	return smtp.SendMail(
		auth.host+":"+auth.port,
		auth.PlainAuth(),
		auth.username,
		mailList,
		[]byte(sub+content),
	)
}
