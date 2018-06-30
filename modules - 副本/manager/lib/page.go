package manager

import (
	"net/http"
	"../../../framework/lib"
	"html/template"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

func fntest1(res http.ResponseWriter, req *http.Request)  {

	// 获取会话
	session, err := xf.Store.Get(req, "g")
	if err != nil {
		res.WriteHeader(404)
		return
	}
	if _, ok := session.Values["user-id"]; !ok {
		http.Redirect(res, req, "/login/index.html", http.StatusFound)
		return
	}
	uid := session.Values["user-id"]

	// 获取数据库链接
	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}
	var user User
	db.Find(&user, uid)

	if user.Id != uid {
		http.Redirect(res, req, "/login/index.html", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles(xf.GetModelesDir("manager")+"/view/index.tpl")
	res.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["user"] = user
	t.Execute(res, data)
}

func pageindex(res http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles(xf.GetModelesDir("manager")+"/view/des.tpl")
	res.Header().Set("Content-Type", "text/html")
	data := make(map[string]string)
	t.Execute(res, data)
}

func page404(res http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles(xf.GetModelesDir("manager")+"/page/404.tpl")
	res.Header().Set("Content-Type", "text/html")
	data := make(map[string]string)
	t.Execute(res, data)
}