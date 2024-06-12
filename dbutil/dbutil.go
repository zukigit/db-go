package dbutil

import (
	_ "github.com/go-sql-driver/mysql"
)

var db Database

func Connect_mysql(dbHost string, dbUser string, dbPasswd string, dbName string, dbPort int, dbTimeoutInSec int) error {

	mysqlDB := NewMysqlDatabase(dbHost, dbUser, dbPasswd, dbName, dbPort, dbTimeoutInSec)

	if err := mysqlDB.Connect(); err != nil {
		return err
	}

	db = mysqlDB
	return nil
}

func Close() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	return db.Close()
}

func Select(unfmt string, arg ...any) ([][]interface{}, error) {
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
