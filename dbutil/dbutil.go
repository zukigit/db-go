package dbutil

import (
	"database/sql"
	"errors"
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

var db *sql.DB

type DBsource struct {
	DBusername string //mandatory
	DBpasswd   string //mandatory
	DBname     string //mandatory
	DBhost     string //optional
	DBport     string //optional
	DBtype     string //no need
}

func ChckDBsource(dbsource DBsource, dbType string) DBsource {
	dbsource.DBtype = dbType
	//check for port and db host
	if dbsource.DBport == "" && dbType == MYSQL {
		dbsource.DBport = "3306"
	}
	if dbsource.DBport == "" && dbType == POSTGRESQL {
		dbsource.DBport = "5432"
	}
	if dbsource.DBhost == "" {
		dbsource.DBhost = "localhost"
	}

	return dbsource
}

func GetDBsource(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) DBsource {
	return DBsource{
		DBusername: DBusername,
		DBpasswd:   DBpasswd,
		DBname:     DBname,
		DBhost:     DBhost,
		DBport:     DBport,
	}
}

func DBconnect(dbsource DBsource) error {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbsource.DBusername, dbsource.DBpasswd, dbsource.DBhost, dbsource.DBport, dbsource.DBname)

	db, err = sql.Open(dbsource.DBtype, dataSourceName)
	if err != nil {
		fmt.Println("Db source is invalid, Error msg: " + err.Error())
		return errors.New(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Can not connect to the databse. Host: " + dbsource.DBhost + ", Error msg: " + pingErr.Error())
		db.Close()
		return errors.New(pingErr.Error())
	}

	fmt.Println("Connected to the db host: " + dbsource.DBhost)
	return nil
}

func DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)
	query := fmt.Sprintf(unfmt, arg...)
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
