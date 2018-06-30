package xf

import (
	"net/http"
	//"github.com/flosch/pongo2"
	"html/template"
	"path/filepath"
	"bytes"
)

func pagination(length int, tn int, format string) template.HTML {
	bt := bytes.NewBuffer([]byte{})
	for n:=0;n<length;n++ {
		t := template.New("")
		t.Parse(format)
		t.Execute(bt, map[string]interface{}{"n":n+1, "enable":n==tn})
	}
	return template.HTML(bt.String())
}

func Display(res http.ResponseWriter, file string, data interface{})  {


	//funcMap := template.FuncMap{"pagination": pagination}
	//t := template.New(file).Funcs(funcMap)
	//t = template.Must(t.ParseFiles("templates/layout.html", "templates/index.html"))
	//t.ExecuteTemplate(res, file, data)

	t1 := template.New(filepath.Base(file))
	t1.Funcs(template.FuncMap{"pagination":pagination})
	t, err := t1.ParseFiles(file)
	//t, err := template.ParseFiles(file)
	if err != nil {
		res.WriteHeader(404)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	t.Execute(res, data)
}

//func Display(res http.ResponseWriter, file string, data interface{})  {
//	var tplExample = pongo2.Must(pongo2.FromFile(file))
//	var err error
//	if data == nil {
//		err = tplExample.ExecuteWriter(pongo2.Context{}, res)
//	} else {
//		err = tplExample.ExecuteWriter(data.(map[string]interface{}).(pongo2.Context{}), res)
//	}
//	if err != nil {
//		http.Error(res, err.Error(), http.StatusInternalServerError)
//	}
//}


