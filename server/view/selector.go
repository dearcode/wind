package view

import (
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/crab/handler"
	"github.com/dearcode/crab/orm"
	"github.com/juju/errors"

	"github.com/dearcode/wind/server"
)

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

	table, ok := tables[r.URL.Query().Get("table")]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rows := table.GetSelector()

	if err = orm.NewStmt(db, s.Table).Query(rows); err != nil {
		log.Errorf("query error:%v, req:%v", errors.ErrorStack(err), r)
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("result rows:%v", rows)

	handler.SendResponseData(w, rows)
}
