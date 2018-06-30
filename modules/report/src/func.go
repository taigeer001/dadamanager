package user

import (
	"fmt"
	"net/http"
	"../../../framework"
	"../../../modules/admin/src"
)

func table_page(res http.ResponseWriter, req *http.Request)  {

	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}

	data := []admin.User{}
	db.Limit(20).Offset(0).Find(&data)
	td := make(map[string]interface{})
	td["list"] = data
	xf.Display(res, xf.GetModelesDir("user/view/table.tpl"), td)
}