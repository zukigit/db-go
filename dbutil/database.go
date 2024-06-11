package dbutil

import (
	"database/sql"
	"fmt"
	"reflect"
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
				var temp_value string
				col_values[i] = &temp_value
			case INT:
				var temp_value int
				col_values[i] = &temp_value
			case UNSIGNED_INT:
				var temp_value uint
				col_values[i] = &temp_value
			case TINYINT:
				var temp_value int8
				col_values[i] = &temp_value
			case UNSIGNED_TINYINT:
				var temp_value uint8
				col_values[i] = &temp_value
			case SMALLINT:
				var temp_value int16
				col_values[i] = &temp_value
			case UNSIGNED_SMALLINT:
				var temp_value uint16
				col_values[i] = &temp_value
			case MEDIUMINT:
				var temp_value int32
				col_values[i] = &temp_value
			case UNSIGNED_MEDIUMINT:
				var temp_value uint32
				col_values[i] = &temp_value
			case BIGINT:
				var temp_value int64
				col_values[i] = &temp_value
			case UNSIGNED_BIGINT:
				var temp_value uint64
				col_values[i] = &temp_value
			case DECIMAL:
				var temp_value float64
				col_values[i] = &temp_value
			case BOOL:
				var temp_value bool
				col_values[i] = &temp_value
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
