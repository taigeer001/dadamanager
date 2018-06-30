package manager

import
(
	"net/http"
	"../../../framework/lib"
)
func Setup()  {
	http.HandleFunc("/login/index.html", DelHandler)
	http.HandleFunc("/login/session", ajaxRegSession)
	http.HandleFunc("/manager/", UseStatic)
	http.HandleFunc("/default/index.html", fntest1)
	http.HandleFunc("/default/des.html", pageindex)
	http.HandleFunc("/errors/404.html", page404)
	xf.BindConnector("user.db", "sqlite3", xf.GetDataDir("user.db"))
}
