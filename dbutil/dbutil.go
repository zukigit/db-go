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

type DButil struct {
	dbUsername string  //mandatory
	dbPasswd   string  //mandatory
	dbName     string  //mandatory
	dbHost     string  //optional
	dbPort     string  //optional
	dbType     string  //no need
	db         *sql.DB //no need
}

func GetInstance(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string, DBtype string) *DButil {
	return &DButil{
		dbUsername: DBusername,
		dbPasswd:   DBpasswd,
		dbName:     DBname,
		dbHost:     DBhost,
		dbPort:     DBport,
		dbType:     DBtype,
	}
}

func (dbsource *DButil) DBconnect() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbsource.dbUsername, dbsource.dbPasswd, dbsource.dbHost, dbsource.dbPort, dbsource.dbName)

	db, err := sql.Open(dbsource.dbType, dataSourceName)
	dbsource.db = db
	if err != nil {
		fmt.Println("Db source is invalid, Error msg: " + err.Error())
		return err
	}

	pingErr := dbsource.db.Ping()
	if pingErr != nil {
		fmt.Println("Can not connect to the databse. Host: " + dbsource.dbHost + ", Error msg: " + pingErr.Error())
		dbsource.db.Close()
		return pingErr
	}

	fmt.Println("Connected to the db host: " + dbsource.dbHost)
	return nil
}

func (dbsource *DButil) DBselect(query string) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)

	rows, err := dbsource.db.Query(query)
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
			case INT, BIGINT:
				var temp_value int64
				col_values[i] = &temp_value
			case DECIMAL:
				var temp_value float32
				col_values[i] = &temp_value
			case BOOL:
				var temp_value bool
				col_values[i] = &temp_value
			}
		}

		err := rows.Scan(col_values...)
		if err != nil {
			fmt.Println("rows scan error ", err)
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

func (dbsource *DButil) DBexec(query string) (int64, error) {
	result, err := dbsource.db.Exec(query)
	if err != nil {
		fmt.Println("db execution error:", err)
		return 0, err
	}

	affected_rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Can not get affected rows, error:", err)
		return 0, err
	}

	return affected_rows, nil
}
