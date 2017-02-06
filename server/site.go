package server

import (
	"github.com/dearcode/petrel/handler"
	"github.com/dearcode/petrel/orm"
	"net/http"
)

type siteInfo struct {
	ID            uint64 `db_default:"auto"`
	Status        bool
	Name          string
	Url           string
	Md5           string
	Level         int
	HotPartBegin  string
	HotPartEnd    string
	ListTagID     uint64
	ClassTagID    uint64
	ContentTagID  uint64
	HtmlFilterID  uint64
	RemoteMysqlID uint64
	Ctime         string `db_defualt:"now()"`
	Mtime         string `db_defualt:"now()"`
}

type site struct {
}

func (s *site) DoGet(w http.ResponseWriter, r *http.Request) {
	dbc := orm.NewDB("127.0.0.1", 3306, "cwind", "root", "guowei", "utf8", "10")
	db, err := dbc.GetConnection()
	if err != nil {
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var infos []siteInfo

	if err = orm.NewStmt(db, "site_info").Query(&infos); err != nil {
		handler.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendResponse(w, http.StatusOK, "%#v", infos)
}
