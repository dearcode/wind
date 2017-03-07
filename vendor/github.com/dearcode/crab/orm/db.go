package orm

import (
	"database/sql"
	"fmt"
	//mysql
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// DB db instance.
type DB struct {
	IP       string
	Port     int
	DBName   string
	UserName string
	Passwd   string
	Charset  string
	TimeOut  string
}

//NewDB create db instance.
func NewDB(ip string, port int, dbName, user, pass, charset, timeout string) *DB {
	return &DB{
		IP:       ip,
		Port:     port,
		DBName:   dbName,
		UserName: user,
		Passwd:   pass,
		Charset:  charset,
		TimeOut:  timeout,
	}
}

// GetConnection open new connect to db.
func (db *DB) GetConnection() (*sql.DB, error) {
	dsn := db.getDSN()
	stmtDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	stmtDB.SetMaxOpenConns(0)
	if err := stmtDB.Ping(); err != nil {
		return nil, err
	}

	return stmtDB, nil
}

func (db *DB) getDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", db.UserName, db.Passwd, db.IP, db.Port)

	if len(db.DBName) > 0 {
		dsn = fmt.Sprintf("%s%s", dsn, db.DBName)

		optStr := db.getOpt()
		if len(optStr) > 0 {
			dsn = fmt.Sprintf("%s?%s", dsn, optStr)
		}
	}

	return dsn
}

func (db *DB) getOpt() string {
	var opts []string
	if len(db.Charset) > 0 {
		opts = append(opts, fmt.Sprintf("charset=%s", db.Charset))
	}

	if len(db.TimeOut) > 0 {
		opts = append(opts, fmt.Sprintf("timeout=%ss", db.TimeOut))
	}

	return strings.Join(opts, "&")
}
