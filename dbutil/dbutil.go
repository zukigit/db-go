package dbutil

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db Database

/*
Purpose: This function creates a connection with mysql and stored it.

Parameters: dataSourceName - it can be generated using Get_mysql_DSN() or mysql.Config{}.FormatDSN().
*/
func Connect_mysql(dataSourceName string) error {
	if db != nil {
		return Err_DB_MULTIPLE_INIT
	}

	mysqlDB := NewMysqlDatabase(dataSourceName)
	if err := mysqlDB.Connect(); err != nil {
		return err
	}

	db = mysqlDB
	return nil
}

/*
Purpose: This function generates the dataSourceName required by Connect_mysql(dataSourceName).

Comment: This function only accepts general parameters such as dbHost, dbUser, dbPasswd, dbName, and dbPort.
For more specific parameters like timeout or others, please use mysql.Config{}.FormatDSN().
Maual link: https://pkg.go.dev/github.com/go-sql-driver/mysql@v1.8.0#Config.FormatDSN
*/
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

/*
Purpose: This function closes the connection and delete it.
*/
func Close() error {
	if db == nil {
		return Err_DB_NOT_INIT
	}
	if err := db.Close(); err != nil {
		return err
	}
	db = nil
	return nil
}

// Purpose: This function executes a SELECT query on the database and returns the result as a matrix (a table of rows and columns).
//
// Return: [][]string: A two-dimensional slice representing the result set as a matrix, where each outer slice represents a row, and each inner slice represents the column values of that row.
// error: An error object if the query execution fails or if the database is not initialized.
//
// Error Handling: If the global db variable is nil, the function returns Err_DB_NOT_INIT indicating that the database has not been initialized.
func Select(unfmt string, arg ...any) ([][]string, error) {
	if db == nil {
		return nil, Err_DB_NOT_INIT
	}
	return db.Select(unfmt, arg...)
}

/*
Purpose: This function returns query's result dynamically.
*/
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
