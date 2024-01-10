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

func DBselect(unfmt string, arg ...any) error {
	query := fmt.Sprintf(unfmt, arg...)
	row_values := make([]map[string]interface{}, 0)

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		raw_col_values := make([]interface{}, len(columns))
		col_values := make(map[string]interface{}, len(columns))

		//preassign empty values in order to compatible with Scan()'s parameter
		for i:= range raw_col_values {
			var temp_value interface{}
			raw_col_values[i] = &temp_value
		}

		err := rows.Scan(raw_col_values...)
		if err != nil {
			return err
		}

		for i, column := range columns {
			// value := reflect.Indirect(reflect.ValueOf(values[i])).Interface()
			value := reflect.Indirect(reflect.ValueOf(raw_col_values[i])).Interface()
			fmt.Println(value)
			col_values[column] = value
		}

		row_values = append(row_values, col_values)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
