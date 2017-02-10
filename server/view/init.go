package view

import (
	"net/http"

	"github.com/dearcode/petrel/handler"
)

func init() {
	handler.Server.AddInterface(&index{}, "/")
	handler.Server.AddInterface(&table{}, "/table/")
	handler.Server.AddHandler(handler.GET, "/static/", true, onStaticGet)
}

//onStaticGet 静态文件
func onStaticGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-control", "no-store")
	http.ServeFile(w, r, "."+r.URL.RequestURI())
}
