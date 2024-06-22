package dbutil

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Stores Database object.
var TEMP_DB Database

// This function needs to be called only once unless you want to change the database configuration again.
// This function stores a MySQL DataSourceName using the provided DataSourceName temporarily.
// You can read how to generate DataSourceName here: [https://pkg.go.dev/github.com/go-sql-driver/mysql@v1.8.0#Config.FormatDSN].
func Init_mysql_DSN(dataSourceName string, maxConCount int) {
	mysqlDB := NewMysqlDatabase(dataSourceName)
	mysqlDB.maxConnections = maxConCount
	TEMP_DB = mysqlDB
}

// This function needs to be called only once unless you want to change the database configuration again.
// It initializes a MySQL DataSourceName using the provided parameters and stores it temporarily.
// The DataSourceName is formatted based on the given host, user, password, database name, and port.
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
// With that [dbutil.Database], You can use [dbutil.Select] or any other utility functions.
// [dbutil.Init_mysql_DSN] or [dbutil.Init_mysql] must be called first to use this function.
func GetConnection() (Database, error) {
	if TEMP_DB == nil {
		return nil, Err_DB_NOT_INIT
	}

	return TEMP_DB, TEMP_DB.Connect()
}

func Close() error {
	return close(TEMP_DB)
}
