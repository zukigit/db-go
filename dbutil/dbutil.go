package dbutil

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db Database

func Connect_mysql(dataSourceName string) error {
	mysqlDB := NewMysqlDatabase(dataSourceName)

	if err := mysqlDB.Connect(); err != nil {
		return err
	}

	db = mysqlDB
	return nil
}

func Get_mysql_DSN(dbHost string, dbUser string, dbPasswd string, dbName string, dbPort int) string {
	if dbPort != 0 {
		dbHost = fmt.Sprintf("%s:%d", dbHost, dbPort)
	}

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPasswd,
		Net:    "tcp",
		Addr:   dbHost,
		DBName: dbName,
	}
	return cfg.FormatDSN()
}

func Close() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	return db.Close()
}

func Select(unfmt string, arg ...any) ([][]string, error) {
	if db == nil {
		return nil, Err_DB_NOT_INIT
	}
	return db.Select(unfmt, arg...)
}

func Execute(unfmt string, arg ...any) (int64, error) {
	if db == nil {
		return 0, Err_DB_NOT_INIT
	}
	return db.Execute(unfmt, arg...)
}

func Begin() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	return db.Begin()
}

func Commit() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	return db.Commit()
}

func Rollback() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	return db.Rollback()
}

// will get failed if you using []interface{} instead of interface{}
func ResultToString(i interface{}) string {
	str, ok := i.(*string)
	if !ok {
		return "FAILED"
	}
	return *str
}
