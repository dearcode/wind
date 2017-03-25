package view

import (
	"html/template"
	"net/http"

	"github.com/dearcode/crab/handler"
	"github.com/dearcode/wind/server"
)

func init() {
	handler.Server.AddInterface(&index{}, "/")
	handler.Server.AddInterface(&table{}, "/table/")
	handler.Server.AddInterface(&item{}, "/item/")
	handler.Server.AddInterface(&detail{}, "/detail/")
	handler.Server.AddInterface(&selector{}, "/selector/")
	handler.Server.AddHandler(handler.GET, "/static/", true, onStaticGet)
}

//onStaticGet 静态文件
func onStaticGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-control", "no-store")
	http.ServeFile(w, r, "."+r.URL.RequestURI())
}

var (
	tables = map[string]server.ViewTable{
		"site": {
			Name:   template.HTML("site"),
			Lable:  "站点信息",
			Object: server.SiteInfo{},

			Fields: []server.ViewField{
				{Name: "ID", Lable: "ID", Widget: server.WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "ListID", Lable: "列表ID"},
				{Name: "List.ID", Lable: "列表ID", Widget: server.WidgetText},
				{Name: "List.Name", Lable: "列表规则", Reference: "list", Relation: "List.ID", Column: "ListID", Widget: server.WidgetSelect, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "Content.ID", Lable: "内容ID", Widget: server.WidgetText},
				{Name: "Content.Name", Lable: "内容规则", Reference: "content", Relation: "Content.ID", Column: "ContentID", Widget: server.WidgetSelect, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "URL", Lable: "URL", Widget: server.WidgetText, Addible: true, Visible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},

		"list": {
			Selector: []server.ListSelector{},
			Object:   server.ListInfo{},
			Name:     template.HTML("list"),
			Lable:    "列表规则",
			Fields: []server.ViewField{
				{Name: "ID", Lable: "ID", Widget: server.WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyBegin", Lable: "内容开始", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyEnd", Lable: "内容结束", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "ItemEnd", Lable: "子项开始", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ItemEnd", Lable: "子项结束", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "URLEnd", Lable: "链接开始", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "URLEnd", Lable: "链接结束", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "PageEnd", Lable: "页签开始", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "PageEnd", Lable: "页签结束", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},

		"content": {
			Selector: []server.ContentSelector{},
			Object:   server.ContentInfo{},
			Name:     template.HTML("content"),
			Lable:    "内容规则",
			Fields: []server.ViewField{
				{Name: "ID", Lable: "ID", Widget: server.WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyBegin", Lable: "内容开始", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyEnd", Lable: "内容结束", Widget: server.WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "DateEnd", Lable: "日期开始", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "DateEnd", Lable: "日期结束", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "TitleEnd", Lable: "标题开始", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "TitleEnd", Lable: "标题结束", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ContentEnd", Lable: "正文开始", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ContentEnd", Lable: "正文结束", Widget: server.WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},
	}
)
