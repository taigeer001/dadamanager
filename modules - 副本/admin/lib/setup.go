package user

import "net/http"

func Setup()  {
	http.HandleFunc("/admin/table.html", fn1)
	http.HandleFunc("/admin/logout", fn2)
	http.HandleFunc("/admin/template.html", fn3)
	http.HandleFunc("/data/table.html", fn1)
	http.HandleFunc("/content/add", fn5)
	http.HandleFunc("/icon/", UseIcon)
}
