package admin

import (
	"../../../framework"
	"net/http"
	"fmt"
)

func login_html(res http.ResponseWriter, req *http.Request)  {
	xf.Display(res, xf.GetModelesDir("admin/view/signin.tpl"), nil)
}

func login_post(res http.ResponseWriter, req *http.Request)  {
	ClearRuquest()
	if !RequestOK(req){
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"请求频繁，稍等5分钟后再试"}))
		return
	}
	// 获取数据库链接
	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}
	// 获取参数
	query := req.URL.Query()
	name := ""; passwd := ""
	if names, ok := query["name"]; ok {
		name = names[0]
	}else {
		res.Write([]byte("args error"))
		return
	}
	if passwds, ok := query["passwd"]; ok {
		passwd = passwds[0]
	}else {
		res.Write([]byte("args error"))
		return
	}
	// 验证
	tmp := User{}
	db.First(&tmp, "Name = ? and Passwd = ?", name, passwd)
	if tmp.Name==name {
		json := xf.ToJson(map[string]interface{}{"type":"success"})
		if !SetSessionUser(&tmp, res, req) {
			res.WriteHeader(404)
			return
		}
		res.Write(json)
	} else {
		RequestUpdate(req)
		res.Write(xf.ToJson(map[string]interface{}{"type":"error", "text":"账号密码错误"}))
	}
}