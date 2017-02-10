package view

import (
	"html/template"
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/petrel/handler"
	"github.com/dearcode/petrel/orm"
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
)

func init() {
	site = server.ViewTable{
		Name:  template.HTML("site"),
		Lable: "站点信息",
		Fields: []server.ViewField{
			{Name: "ID", Lable: "ID", Widget: server.WidgetText},
			{Name: "Name", Lable: "名称", Widget: server.WidgetText, Sortable: true, Addible: true, Visible: true, Modifiable: true},
			{Name: "URL", Lable: "URL", Widget: server.WidgetText, Sortable: false, Addible: true, Visible: true, Modifiable: true},
			{Name: "Mtime", Lable: "更新时间", Widget: server.WidgetText, Sortable: true, Addible: false, Visible: true, Modifiable: true, Readonly: true},
		},
	}
}

type index struct {
}

func (i *index) DoGet(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./html/index.html", "./html/navbar.html", "./html/footer.html", "./html/header.html", "./html/common.html"))
	if err := t.Execute(w, site); err != nil {
		log.Errorf("Execute error:%v, site:%+v", err, site)
		return
	}
	log.Debugf("site:%+v", site)
}
