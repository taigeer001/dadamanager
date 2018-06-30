package user

import (
	"net/http"
)

func Setup() {




	http.HandleFunc("/user/table", table_page)
	http.HandleFunc("/user/add.html", table_add)


}