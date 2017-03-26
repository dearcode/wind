package view

import (
	"html/template"
	"net/http"

	"github.com/davygeek/log"

	"github.com/dearcode/wind/server/function"
)

type detail struct {
}

//DoGet 详情页面
func (d *detail) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("detail.html").Funcs(function.Common).ParseFiles(
		"./html/detail.html",
		"./html/navbar.html",
		"./html/footer.html",
		"./html/header.html",
		"./html/edit.html",
	))
	t = t.Funcs(function.Common)
	table, ok := tables[r.URL.Query().Get("table")]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	table.ID = template.JS(r.URL.Query().Get("id"))

	if err := t.Execute(w, table); err != nil {
		log.Errorf("Execute error:%v, table:%+v", err, table)
		return
	}
	log.Debugf("site:%+v", table)
}

type index struct {
}

//DoGet 首页
func (i *index) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index.html").Funcs(function.Common).ParseFiles(
		"./html/index.html",
		"./html/navbar.html",
		"./html/footer.html",
		"./html/header.html",
		"./html/list.html",
		"./html/edit.html",
		"./html/message.html",
		"./html/common.html",
	))

	name := r.URL.Query().Get("table")
	if name == "" {
		name = "site"
	}

	table, ok := tables[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := t.Execute(w, table); err != nil {
		log.Errorf("Execute error:%v, site:%+v", err, table)
		return
	}
	log.Debugf("site:%+v", table)
}
