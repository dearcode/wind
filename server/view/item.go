package view

import (
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/crab/handler"
	"github.com/dearcode/crab/orm"
	"github.com/juju/errors"

	"github.com/dearcode/wind/server"
)

type item struct {
	ID    int64  `json:"id"`
	Table string `json:"table"`
}

//DoGet 获取单个项所有字段
func (i *item) DoGet(w http.ResponseWriter, r *http.Request) {
	if err := handler.ParseURLVars(r, i); err != nil {
		log.Errorf("invalid request:%v, error:%v", r, err)
		handler.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	i.Table = orm.FieldEscape(i.Table)

	db, err := server.DBC.GetConnection()
	if err != nil {
		log.Errorf("GetConnection error:%v, req:%v", err, r)
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	table, ok := tables[i.Table]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := table.GetObject()

	if err = orm.NewStmt(db, i.Table).Query(data); err != nil {
		log.Errorf("query error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("result data:%v", data)

	handler.SendResponseData(w, data)
}
