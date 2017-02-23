package server

import (
	"github.com/dearcode/petrel/orm"
)

var (
	//DBC db connect.
	DBC *orm.DB
)

func init() {
	DBC = orm.NewDB("127.0.0.1", 3306, "cwind", "root", "guowei", "utf8", "10")
}
