package dbutil

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Database interface {
	Select(unfmt string, arg ...any) ([][]interface{}, error)
	Execute(unfmt string, arg ...any) (int64, error)
	Begin() error
	Commit() error
	Rollback() error
	Close() error
}

type MysqlDatabase struct {
	db        *sql.DB
	err       *error
	isInTranx *bool
}

func (mysql MysqlDatabase) dbSelect(query string) ([][]interface{}, error) {
	row_values := make([][]interface{}, 0)

	rows, err := mysql.db.Query(query)
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

func (mysql MysqlDatabase) Select(unfmt string, arg ...any) ([][]interface{}, error) {
	if isDBinit() {
		query := fmt.Sprintf(unfmt, arg...)
		return mysql.dbSelect(query)
	} else {
		err = Err_DB_NOT_INIT
	}
	return nil, err
}

func (mysql MysqlDatabase) Close() error {
	return mysql.db.Close()
}
