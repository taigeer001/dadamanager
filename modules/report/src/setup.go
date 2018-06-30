package user

import (
	"net/http"
)

func Setup() {

	http.HandleFunc("/report/table", table_page)


}