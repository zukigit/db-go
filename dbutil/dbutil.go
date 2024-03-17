package dbutil

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"
)

const (
	VARCHAR  = "VARCHAR"
	TEXT     = "TEXT"
	NVARCHAR = "NVARCHAR"
	DECIMAL  = "DECIMAL"
	BOOL     = "BOOL"
	INT      = "INT"
	BIGINT   = "BIGINT"
)

var db *sql.DB //no need
var err error

func dbPing() error {
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func dbConnect(dbType string, dataSourceName string) error {
	db, err = sql.Open(dbType, dataSourceName)
	if err != nil {
		return err
	}

	return dbPing()
}

func Connect_mysql(dbHost string, dbUser string, dbPasswd string, dbName string, dbPort int, dbTimeoutInSec int) error {
	dbType := MYSQL
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == 0 {
		dbPort = 3306
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds",
		dbUser, dbPasswd, dbHost, dbPort, dbName, dbTimeoutInSec)

	return dbConnect(dbType, dataSourceName)
}

func Close() error {
	return db.Close()
}

func Select(query string) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)

	rows, err := db.Query(query)
	if err != nil {
		return row_values, err
	}
	defer rows.Close()

	columns, err := rows.ColumnTypes()
	if err != nil {
		return row_values, err
	}

	for rows.Next() {
		col_values := make([]interface{}, len(columns))

		//defines col_values' types
		for i := range col_values {
			switch columns[i].DatabaseTypeName() {
			case VARCHAR, NVARCHAR, TEXT:
				var temp_value string
				col_values[i] = &temp_value
			case INT:
				var temp_value int
				col_values[i] = &temp_value
			case BIGINT:
				var temp_value int64
				col_values[i] = &temp_value
			case DECIMAL:
				var temp_value float64
				col_values[i] = &temp_value
			case BOOL:
				var temp_value bool
				col_values[i] = &temp_value
			}
		}

		err := rows.Scan(col_values...)
		if err != nil {
			return row_values, err
		}

		//parses to readable data types from interface
		for i := range col_values {
			col_values[i] = reflect.Indirect(reflect.ValueOf(col_values[i])).Interface()
		}

		row_values = append(row_values, col_values)
	}
	return row_values, nil
}

func DBexec(query string) (int64, error) {
	result, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	affected_rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected_rows, err
}
