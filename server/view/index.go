package view

import (
	"html/template"
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/crab/handler"
	"github.com/dearcode/crab/orm"
	"github.com/juju/errors"

	"github.com/dearcode/wind/server"
)

type table struct {
	Model  string `json:"model"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type selector struct {
	Table string `json:"table"`
}

func (s *selector) DoGet(w http.ResponseWriter, r *http.Request) {
	if err := handler.ParseURLVars(r, s); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	s.Table = orm.FieldEscape(s.Table)

	db, err := server.DBC.GetConnection()
	if err != nil {
		log.Errorf("GetConnection error:%v, req:%v", err, r)
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	var rows interface{}

	switch s.Table {
	case "list":
		rows = &[]server.ListSelector{}
	}

	if err = orm.NewStmt(db, s.Table).Query(rows); err != nil {
		log.Errorf("query error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("result rows:%v", rows)

	handler.SendResponseData(w, rows)
}

func (t *table) DoPost(w http.ResponseWriter, r *http.Request) {
	if err := handler.ParseURLVars(r, t); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}

	switch t.Model {
	case "site":
		data = &server.SiteInfo{}
	}

	if err := handler.ParseFormVars(r, data); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := server.DBC.GetConnection()
	if err != nil {
		log.Errorf("GetConnection error:%v, req:%v", err, r)
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	id, err := orm.NewStmt(db, t.Model).Insert(data)
	if err != nil {
		log.Errorf("insert:%v error:%v", data, errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("insert %v, id:%v", data, id)

	handler.SendResponseData(w, nil)
}

func (t *table) DoGet(w http.ResponseWriter, r *http.Request) {
	if err := handler.ParseURLVars(r, t); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := server.DBC.GetConnection()
	if err != nil {
		log.Errorf("GetConnection error:%v, req:%v", err, r)
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	stmt := orm.NewStmt(db, t.Model).Sort(t.Sort).Order(t.Order).Offset(t.Offset).Limit(t.Limit)

	total, err := stmt.Count()
	if err != nil {
		log.Errorf("count error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if total < 1 {
		log.Infof("req:%v, data not found", r)
		handler.SendRows(w, total, &server.SiteInfo{})
		return
	}

	var rows interface{}

	switch t.Model {
	case "site":
		rows = &[]server.SiteInfo{}
	}

	if err = stmt.Query(rows); err != nil {
		log.Errorf("query error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("site count:%d, rows:%v", total, rows)

	handler.SendRows(w, total, rows)
}

var (
	site server.ViewTable
	list server.ViewTable
)

func init() {
	site = server.ViewTable{
		Name:  template.HTML("site"),
		Lable: "站点信息",
		Fields: []server.ViewField{
			{Name: "ID", Lable: "ID", Widget: server.WidgetText, Readonly: true},
			{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
			{Name: "List.ID", Lable: "列表ID", Widget: server.WidgetText},
			{Name: "List.Name", Lable: "列表", Reference: "List", Relation: "List.ID", Column: "ListID", Widget: server.WidgetSelect, Sortable: true, Addible: true, Visible: true, Modifiable: true},
			{Name: "URL", Lable: "URL", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Addible: false, Visible: true, Modifiable: true, Readonly: true},
		},
	}

	list = server.ViewTable{
		Name:  template.HTML("list"),
		Lable: "列表规则",
		Fields: []server.ViewField{
			{Name: "ID", Lable: "ID", Widget: server.WidgetText, Readonly: true},
			{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
			{Name: "BodyBegin", Lable: "内容开始", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "BodyEnd", Lable: "内容结束", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "URL", Lable: "URL", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "URL", Lable: "URL", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "URL", Lable: "URL", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Addible: false, Visible: true, Modifiable: true, Readonly: true},
		},
	}

}

type index struct {
}

func (i *index) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./html/index.html", "./html/navbar.html", "./html/footer.html", "./html/header.html", "./html/common.html"))
	var table server.ViewTable

	switch r.URL.Query().Get("table") {
	default:
		table = site
	case "List":
		table = list
	}

	if err := t.Execute(w, table); err != nil {
		log.Errorf("Execute error:%v, site:%+v", err, site)
		return
	}
	log.Debugf("site:%+v", site)
}
