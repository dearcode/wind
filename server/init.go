package server

import (
	"github.com/dearcode/crab/config"
	"github.com/dearcode/crab/orm"
)

var (
	//DBC db connect.
	DBC *orm.DB
	//Conf 通用配置文件.
	Conf windConfig
)

func init() {

}

type windConfig struct {
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

//Init 连接数据库.
func Init() error {
	if err := config.LoadConfig("conf/wind.ini", &Conf); err != nil {
		return err
	}

	DBC = orm.NewDB(Conf.DB.Host, Conf.DB.Port, Conf.DB.Name, Conf.DB.User, Conf.DB.Password, "utf8", "10")
	db, err := DBC.GetConnection()
	if err != nil {
		return err
	}
	db.Close()

	return nil
}
