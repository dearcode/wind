package view

import (
	"html/template"
	"net/http"

	"github.com/dearcode/crab/handler"
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
	tables = map[string]Table{
		"site": {
			Name:   template.HTML("site"),
			Lable:  "站点信息",
			Object: SiteInfo{},
			Fields: []Field{
				{Name: "ID", Lable: "ID", Widget: WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "ListID", Lable: "列表ID"},
				{Name: "List.ID", Lable: "列表ID", Widget: WidgetText},
				{Name: "List.Name", Lable: "列表规则", Relation: "List.ID", Widget: WidgetSelect, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "Content.ID", Lable: "内容ID", Widget: WidgetText},
				{Name: "Content.Name", Lable: "内容规则", Relation: "Content.ID", Widget: WidgetSelect, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "Filter.ID", Lable: "替换ID", Widget: WidgetText},
				{Name: "Filter.Name", Lable: "过滤替换", Relation: "Filter.ID", Widget: WidgetSelectMore, Addible: true, Visible: true, Modifiable: true},
				{Name: "URL", Lable: "URL", Widget: WidgetText, Addible: true, Visible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},

		"list": {
			Object: ListInfo{},
			Name:   template.HTML("list"),
			Lable:  "列表规则",
			Fields: []Field{
				{Name: "ID", Lable: "ID", Widget: WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyBegin", Lable: "内容开始", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyEnd", Lable: "内容结束", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "ItemEnd", Lable: "子项开始", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ItemEnd", Lable: "子项结束", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "URLEnd", Lable: "链接开始", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "URLEnd", Lable: "链接结束", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "PageEnd", Lable: "页签开始", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "PageEnd", Lable: "页签结束", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},

		"content": {
			Object: ContentInfo{},
			Name:   template.HTML("content"),
			Lable:  "内容规则",
			Fields: []Field{
				{Name: "ID", Lable: "ID", Widget: WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyBegin", Lable: "内容开始", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "BodyEnd", Lable: "内容结束", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "DateEnd", Lable: "日期开始", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "DateEnd", Lable: "日期结束", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "TitleEnd", Lable: "标题开始", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "TitleEnd", Lable: "标题结束", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ContentEnd", Lable: "正文开始", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "ContentEnd", Lable: "正文结束", Widget: WidgetTextArea, Addible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},
		"filter": {
			Object: FilterInfo{},
			Name:   template.HTML("filter"),
			Lable:  "过滤替换",
			Fields: []Field{
				{Name: "ID", Lable: "ID", Widget: WidgetText, Readonly: true},
				{Name: "Name", Lable: "名称", Widget: WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
				{Name: "Key1", Lable: "内容1", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "Key2", Lable: "内容2", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "Type", Lable: "类型", Widget: WidgetTextArea, Addible: true, Visible: true, Modifiable: true},
				{Name: "Mtime", Lable: "更新时间", Widget: WidgetText, Sortable: true, Visible: true, Modifiable: true, Readonly: true},
			},
		},
	}
)
