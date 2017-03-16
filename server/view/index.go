package view

import (
	"html/template"
	"net/http"

	"github.com/davygeek/log"

	"github.com/dearcode/wind/server"
)

type detail struct {
}

//DoGet 详情页面
func (d *detail) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./html/detail.html",
		"./html/navbar.html",
		"./html/footer.html",
		"./html/header.html",
		"./html/edit.html",
	))
	var table server.ViewTable

	switch r.URL.Query().Get("table") {
	case "list", "List":
		table = list
	case "content":
		table = content
	default:
		table = site
	}

	table.ID = template.JS(r.URL.Query().Get("id"))

	if err := t.Execute(w, table); err != nil {
		log.Errorf("Execute error:%v, site:%+v", err, site)
		return
	}
	log.Debugf("site:%+v", site)
}

type index struct {
}

//DoGet 首页
func (i *index) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./html/index.html",
		"./html/navbar.html",
		"./html/footer.html",
		"./html/header.html",
		"./html/list.html",
		"./html/edit.html",
		"./html/message.html",
		"./html/common.html",
	))
	var table server.ViewTable

	switch r.URL.Query().Get("table") {
	case "list":
		table = list
	case "content":
		table = content
	default:
		table = site
	}

	if err := t.Execute(w, table); err != nil {
		log.Errorf("Execute error:%v, site:%+v", err, site)
		return
	}
	log.Debugf("site:%+v", site)
}
