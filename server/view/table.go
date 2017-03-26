package view

import (
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/crab/handler"
	"github.com/dearcode/crab/orm"
	"github.com/juju/errors"

	"github.com/dearcode/wind/server"
)

type table struct {
	ID     int64  `json:"id"`
	Model  string `json:"model"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (t *table) DoPut(w http.ResponseWriter, r *http.Request) {
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

	table, ok := tables[t.Model]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := table.GetObject()

	if err := orm.NewStmt(db, t.Model).Where("%s.id=%d", t.Model, t.ID).Query(data); err != nil {
		log.Errorf("query error:%v, req:%+v", err, r)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.ParseFormVars(r, data); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := orm.NewStmt(db, t.Model).Where("%s.id=%d", t.Model, t.ID).Update(data)
	if err != nil {
		log.Errorf("update:%v error:%v", data, errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("update:%v, id:%v", data, id)
	handler.SendResponseData(w, nil)
}

func (t *table) DoPost(w http.ResponseWriter, r *http.Request) {
	if err := handler.ParseURLVars(r, t); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	table, ok := tables[t.Model]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := table.GetObject()

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
		handler.SendRows(w, total, &SiteInfo{})
		return
	}

	table, ok := tables[t.Model]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	rows := table.GetObjectSlice()

	if err = stmt.Query(rows); err != nil {
		log.Errorf("query error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("site count:%d, rows:%v", total, rows)

	handler.SendRows(w, total, rows)
}
