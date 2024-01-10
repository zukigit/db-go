package dbutil

import (
	"database/sql"
	"errors"
	"fmt"

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

func DBselect(unfmt string, arg ...any) error {
	columns_count := 0
	query := fmt.Sprintf(unfmt, arg...)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	col_types, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	columns_count = len(col_types)

	for rows.Next() {
		columns_values := make([]interface{}, columns_count)

		for i := range columns_values {
			var temp_value interface{}
			switch col_types[i].DatabaseTypeName() {
			case VARCHAR:
			case TEXT:
			case NVARCHAR:
				temp_value = new(string)
			case DECIMAL:
			case INT:
			case BIGINT:
				temp_value = new(int)
			case BOOL:
				temp_value = new(bool)
			}
			columns_values[i] = &temp_value
		}
		err := rows.Scan(columns_values...)
		if err != nil {
			fmt.Println("error at scanning", err)
			return err
		}
		for i, colValue := range columns_values {
			switch v := colValue.(type) {
			case *string:
				// Handle string value
				fmt.Printf("Column %d: %s\n", i, *v)
			case *int:
				// Handle int value
				fmt.Printf("Column %d: %d\n", i, *v)
			case *bool:
				// Handle bool value
				fmt.Printf("Column %d: %t\n", i, *v)
			// Add cases for other types as needed
			default:
				fmt.Println("type:", v)
				fmt.Printf("Column %d: Unexpected type\n", i)
			}
		}
	}

	return nil
}
