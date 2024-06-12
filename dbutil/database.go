package dbutil

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Database interface {
	Connect() error
	Ping() error
	Select(unfmt string, arg ...any) ([][]interface{}, error)
	Execute(unfmt string, arg ...any) (int64, error)
	Begin() error
	Commit() error
	Rollback() error
	Close() error
}

func dbSelect(query string, db *sql.DB) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		col_values := make([]interface{}, len(columns))

		//defines col_values' types
		for i := range col_values {
			switch columns[i].DatabaseTypeName() {
			case VARCHAR, NVARCHAR, TEXT:
				col_values[i] = new(*string)
			case INT:
				col_values[i] = new(*int)
			case UNSIGNED_INT:
				col_values[i] = new(*uint)
			case TINYINT:
				col_values[i] = new(*int8)
			case UNSIGNED_TINYINT:
				col_values[i] = new(*uint8)
			case SMALLINT:
				col_values[i] = new(*int16)
			case UNSIGNED_SMALLINT:
				col_values[i] = new(*uint16)
			case MEDIUMINT:
				col_values[i] = new(*int32)
			case UNSIGNED_MEDIUMINT:
				col_values[i] = new(*uint32)
			case BIGINT:
				col_values[i] = new(*int64)
			case UNSIGNED_BIGINT:
				col_values[i] = new(*uint64)
			case DECIMAL:
				col_values[i] = new(*float64)
			case BOOL:
				col_values[i] = new(*bool)
			case DATE, DATETIME, TIMESTAMP, TIME, YEAR:
				col_values[i] = new(*time.Time)
			case BINARY, VARBINARY, BLOB, MEDIUMBLOB, LONGBLOB:
				col_values[i] = new(*string)
			case POINT, GEOMETRY:
				col_values[i] = new(*string)
			default:
				fmt.Println("Column type:", columns[i].DatabaseTypeName())
				return nil, Err_UNDEFINED_COLLUMN_TYPE
			}
		}

		err := rows.Scan(col_values...)
		if err != nil {
			return nil, err
		}

		//parses memory addresses to values
		for i := range col_values {
			col_values[i] = reflect.Indirect(reflect.ValueOf(col_values[i])).Interface()
		}

		row_values = append(row_values, col_values)
	}
	return row_values, nil
}

func dbExecute(query string, db *sql.DB) (int64, error) {
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

func dbBegin(query string, db *sql.DB) error {
	_, err := dbExecute(query, db)
	if err != nil {
		return err
	}
	return nil
}

func dbCommit(query string, db *sql.DB) error {
	_, err := dbExecute(query, db)
	if err != nil {
		return err
	}
	return nil
}

func dbRollback(query string, db *sql.DB) error {
	_, err := dbExecute(query, db)
	if err != nil {
		return err
	}
	return nil
}
