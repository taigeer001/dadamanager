package user

import (
	"fmt"
	"net/http"
	"../../../framework"
	"../../../modules/admin/src"
	"strconv"
)

func table_page(res http.ResponseWriter, req *http.Request)  {

	user := admin.GetSessionUser(req)

	if user == nil {
		res.WriteHeader(404)
		return
	}

	if !user.Enforce("user", "read") {
		res.WriteHeader(404)
		return
	}

	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}

	// 获取参数
	query := req.URL.Query()
	pn := 1
	if names, ok := query["n"]; ok {
		pn, err = strconv.Atoi(names[0])
		if err != nil {
			pn = 1
		}
	}

	pl := 1
	data := []admin.User{}; count:=0
	db.Limit(pl).Offset(pl*(pn-1)).Find(&data)
	db.Table("users").Count(&count)
	length := count/pl
	if length <= 0 {
		length = 1
	}
	td := map[string]interface{}{"list":data, "count": count, "page": length, "n": pn}
	xf.Display(res, xf.GetModelesDir("user/view/table.tpl"), td)
}

func table_add(res http.ResponseWriter, req *http.Request)  {
	xf.Display(res, xf.GetModelesDir("user/view/add.tpl"), nil)
}
