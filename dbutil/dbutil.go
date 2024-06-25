package dbutil

import (
	"fmt"
	"sync"

	"github.com/go-sql-driver/mysql"
)

// Stores Database object.
var TEMP_DB Database

var MAX_CONS int
var CURRENT_CONS int
var MUTEX *sync.Mutex

// It initializes MysqlDatabase using the provided dataSourceName and stores it.
//
//	dataSourceName (string): The DSN string that contains the information required to connect to the MySQL database.
//	maxConCount (int): The maximum number of connections to allow in the connection pool. If set to 0, it allows for an unlimited number of connections.
//
// You can read how to generate DataSourceName here: [https://pkg.go.dev/github.com/go-sql-driver/mysql@v1.8.0#Config.FormatDSN].
func Init_mysql_DSN(dataSourceName string, maxConCount int) {
	mysqlDB := NewMysqlDatabase(dataSourceName)
	MAX_CONS = maxConCount
	TEMP_DB = mysqlDB
	MUTEX = &sync.Mutex{}
}

// It initializes MysqlDatabase using the provided parameters and stores it.
//
//	maxConCount (int): The maximum number of connections to allow in the connection pool. If set to 0, it allows for an unlimited number of connections.
//
// This function needs to be called only once unless you want to change the database configuration again.
// If you want to use your own DataSourceName, you can use [dbutil.Init_mysql_DSN] instead.
func Init_mysql(dbHost string, dbUser string, dbPasswd string, dbName string, dbPort int, maxConCount int) {
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

	Init_mysql_DSN(cfg.FormatDSN(), maxConCount)
}

// This function establishes a connection to the database and returns [dbutil.Database].
// To use this function, [dbutil.Init_mysql_DSN] or [dbutil.Init_mysql] must be called first.
func GetCon() (Database, error) {
	if TEMP_DB == nil {
		return nil, Err_DB_NOT_INIT
	}

	return TEMP_DB.Connect()
}

func Close() error {
	return close(TEMP_DB)
}
