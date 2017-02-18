package server

import (
	"github.com/dearcode/petrel/orm"
)

var (
	//DBC db connect.
	DBC *orm.DB
)

func init() {
	DBC = orm.NewDB("192.168.199.199", 3306, "cwind", "orm_test", "orm_test_password", "utf8", "10")
}
