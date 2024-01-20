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

var db *sql.DB

type DBsource struct {
	dbUsername string  //mandatory
	dbPasswd   string  //mandatory
	dbName     string  //mandatory
	dbHost     string  //optional
	dbPort     string  //optional
	dbType     string  //no need
	db         *sql.DB //no need
}

func ChckDBsource(dbsource DBsource) DBsource {
	//check for port and db host
	if dbsource.dbPort == "" && dbsource.dbType == MYSQL {
		dbsource.dbPort = "3306"
	}
	if dbsource.dbPort == "" && dbsource.dbType == POSTGRESQL {
		dbsource.dbPort = "5432"
	}
	if dbsource.dbHost == "" {
		dbsource.dbHost = "localhost"
	}

	return dbsource
}

func GetDBsource(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string, DBtype string) DBsource {
	return DBsource{
		dbUsername: DBusername,
		dbPasswd:   DBpasswd,
		dbName:     DBname,
		dbHost:     DBhost,
		dbPort:     DBport,
		dbType:     DBtype,
	}
}

func (dbsource *DBsource) DBconnect() error {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbsource.dbUsername, dbsource.dbPasswd, dbsource.dbHost, dbsource.dbPort, dbsource.dbName)

	dbsource.db, err = sql.Open(dbsource.dbType, dataSourceName)
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

func (dbsource *DBsource) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)
	query := fmt.Sprintf(unfmt, arg...)

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
