package admin

import (
	"../../../framework"
	"net/http"
)

func index_page(res http.ResponseWriter, req *http.Request)  {
	user := GetSessionUser(req)
	if user == nil {
		http.Redirect(res, req, "/admin/signin.html", http.StatusFound)
		return
	}
	data := map[string]interface{}{"user":user}
	xf.Display(res, xf.GetModelesDir("admin/view/index.tpl"), data)
}

func sginout(res http.ResponseWriter, req *http.Request)  {
	ClearSessionUser(res, req)
	res.Write(xf.ToJson(map[string]interface{}{"type":"success"}))
}