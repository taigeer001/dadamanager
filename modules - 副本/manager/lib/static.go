package manager

import (
	"net/http"
	"mime"
	"path/filepath"
	"os"
	"io"
	"../../../framework/lib"
)

func UseStatic(res http.ResponseWriter, req *http.Request)  {
	ctype := mime.TypeByExtension(filepath.Ext(req.URL.Path[8:]))
	src, err := os.Open(xf.GetModelesDir("manager")+"/static/"+req.URL.Path[8:])
	if err != nil {
		res.WriteHeader(404)
		return
	}
	res.Header().Set("Content-Type", ctype)
	io.Copy(res, src)
}

