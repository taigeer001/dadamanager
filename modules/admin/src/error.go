package admin

import (
	"../../../framework"
	"net/http"
)

func page_404(res http.ResponseWriter, req *http.Request)  {
	xf.Display(res, xf.GetModelesDir("admin/res/404.tpl"), nil)
}