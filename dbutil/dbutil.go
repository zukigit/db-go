package dbutil

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var TEMP_DB Database

func Init_mysql(dbHost string, dbUser string, dbPasswd string, dbName string, dbPort int) {
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

	TEMP_DB = NewMysqlDatabase(cfg.FormatDSN())
}

func Connect() (Database, error) {
	if TEMP_DB == nil {
		return nil, Err_DB_NOT_INIT
	}

	return TEMP_DB.Connect()
}
