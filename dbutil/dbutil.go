package dbutil

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Store database object temporarily
var TEMP_DB Database

// This function initializes a MySQL database connection using the provided data source name and stores it temporarily.
// You can read how to generate DataSourceName [here]: https://pkg.go.dev/github.com/go-sql-driver/mysql@v1.8.0#Config.FormatDSN
func Init_mysql_DSN(dsn string) {
	TEMP_DB = NewMysqlDatabase(dsn)
}

// This function initializes a MySQL database connection using the provided parameters and stores it temporarily.
// It formats the connection string based on the provided host, user, password, database name, and port.
// If you want to use DataSourceName by yourself, you can use [dbutil.Init_mysql_DSN] instead.
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
