package admin

import (
	"net/http"
	"fmt"
	"../../../framework"
)

func test(res http.ResponseWriter, req *http.Request)  {
	user := GetUser(1)
	fmt.Println(user)
	SetSessionUser(user, res, req)
}

func test1(res http.ResponseWriter, req *http.Request)  {
	fmt.Println(GetSessionUser(req))
}

func test2(res http.ResponseWriter, req *http.Request)  {

	ClearSessionUser(res, req)
}

func Setup() {

	http.HandleFunc("/admin/test", test)
	http.HandleFunc("/admin/test1", test1)
	http.HandleFunc("/admin/test2", test2)

	http.HandleFunc("/admin/signin.html", login_html)
	http.HandleFunc("/admin/signin", login_post)
	http.HandleFunc("/admin/forgot.html", forgot_passwd)
	http.HandleFunc("/admin/forgot", send_email)
	http.HandleFunc("/admin/index.html", index_page)
	http.HandleFunc("/admin/signout", sginout)

	http.HandleFunc("/error/404.html", page_404)

	xf.BindConnector("user.db", "sqlite3", xf.GetDataDir("database/user.db"))
	xf.HandleStatic("/admin/static/", xf.GetModelesDir("admin/static"))
}