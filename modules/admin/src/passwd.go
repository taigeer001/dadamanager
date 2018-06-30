package admin

import (
	"../../../framework"
	"net/http"
	"../../../modules/smtp/src"
)

func forgot_passwd(res http.ResponseWriter, req *http.Request)  {
	xf.Display(res, xf.GetModelesDir("admin/view/forgot.tpl"), nil)
}

func send_email(res http.ResponseWriter, req *http.Request)  {
	ClearRuquest()
	if !RequestDomianOK("forgot", req) {
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"请求频繁，稍等5分钟后再试"}))
		return
	}
	query := req.URL.Query()
	mail := ""; name := ""
	if names, ok := query["mail"]; ok {
		mail = names[0]
	}else {
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"外部错误"}))
		return
	}
	if names, ok := query["name"]; ok {
		name = names[0]
	}else {
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"外部错误"}))
		return
	}
	// 获取数据库链接
	db, err := xf.GetConnect("user.db")
	if err != nil {
		res.WriteHeader(404)
		return
	}
	var user User
	db.Where("Name = ? and Email = ?", name, mail).First(&user)
	if user.Id == 0 {
		RequestDomianUpdate("forgot", req)
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"邮箱错误"}))
		return
	}
	if smtp.SendMail(mail, "找回密码", user.Passwd) == nil {
		res.Write(xf.ToJson(map[string]interface{}{"type":"success"}))
	} else {
		RequestDomianUpdate("forgot", req)
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"邮件发送错误"}))
	}
}