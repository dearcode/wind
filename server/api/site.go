package api

import (
	"net/http"

	"github.com/davygeek/log"
	"github.com/dearcode/petrel/handler"
	"github.com/dearcode/petrel/orm"
	"github.com/juju/errors"

	"github.com/dearcode/wind/server"
)

type site struct {
}

type sites struct {
	Total int64
	Rows  []server.SiteInfo
}

func (ss *sites) DoGet(w http.ResponseWriter, r *http.Request) {
	ss = &sites{}
	db, err := server.DBC.GetConnection()
	if err != nil {
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	if ss.Total, err = orm.NewStmt(db, "site_info").Count(); err != nil {
		log.Errorf("count error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = orm.NewStmt(db, "site_info").Query(&ss.Rows); err != nil {
		log.Errorf("query error:%v", errors.ErrorStack(err))
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Debugf("site rows:%v", ss)

	handler.SendRows(w, ss.Total, ss.Rows)
}
