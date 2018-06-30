package user

import (
	"net/http"
	"../../../framework/lib"
	"../../../modules/manager/lib"
	"html/template"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"encoding/json"
	"mime"
	"path/filepath"
	"os"
	"io"
)

func fn1(res http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles(xf.GetModelesDir("admin")+"/view/table.tpl")
	res.Header().Set("Content-Type", "text/html")

	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}
	data := []manager.User{}
	db.Limit(20).Offset(0).Find(&data)
	td := make(map[string]interface{})
	td["list"] = data
	t.Execute(res, data)
}


func fn2(res http.ResponseWriter, req *http.Request)  {

	session, err := xf.Store.Get(req, "g")
	if err != nil {
		res.WriteHeader(404)
		return
	}
	delete(session.Values, "user-id")
	session.Save(req, res)

	b, err := json.Marshal(map[string]interface{}{"type":"sucess"})
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	res.Write(b)

}


func fn3(res http.ResponseWriter, req *http.Request)  {

	t, _ := template.ParseFiles(xf.GetModelesDir("admin")+"/view/template.tpl")
	res.Header().Set("Content-Type", "text/html")

	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}
	data := []Template{}
	db.Limit(20).Offset(0).Find(&data)
	td := make(map[string]interface{})
	td["list"] = data
	t.Execute(res, data)

}


func fn5(res http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles(xf.GetModelesDir("admin")+"/view/add.tpl")
	res.Header().Set("Content-Type", "text/html")

	db, err := xf.GetConnect("user.db")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		return
	}
	data := []manager.User{}
	db.Limit(20).Offset(0).Find(&data)
	td := make(map[string]interface{})
	td["list"] = data
	t.Execute(res, data)
}

func UseIcon(res http.ResponseWriter, req *http.Request)  {
	ctype := mime.TypeByExtension(filepath.Ext(req.URL.Path[5:]))
	src, err := os.Open(xf.GetDataDir("template")+"/"+req.URL.Path[5:])
	if err != nil {
		res.WriteHeader(404)
		return
	}
	res.Header().Set("Content-Type", ctype)
	io.Copy(res, src)
}
