package manager

import (
	"net/http"
	"../../../framework/lib"
	"html/template"
	"fmt"
	"encoding/json"
)
func DelHandler(res http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles(xf.GetModelesDir("manager")+"/view/signin.tpl")
	res.Header().Set("Content-Type", "text/html")
	data := make(map[string]string)
	t.Execute(res, data)
}

func ajaxRegSession(res http.ResponseWriter, req *http.Request)  {
	// 获取会话
	session, err := xf.Store.Get(req, "g")
	if err != nil {
		res.WriteHeader(404)
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
	db.First(&tmp, "username = ? and passwd = ?", name, passwd)
	if tmp.Name==name {
		b, err := json.Marshal(map[string]interface{}{"type":"sucess"})
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}
		session.Values["user-id"] = tmp.Id
		session.Save(req, res)
		res.Write(b)
	} else {
		b, err := json.Marshal(map[string]interface{}{"type":"error", "text":"账号密码错误"})
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}
		res.Write(b)
	}
	defer db.Close()
}