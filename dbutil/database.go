package dbutil

import (
	"database/sql"
)

type Database interface {
	Connect() error
	Ping() error
	Select(unfmt string, arg ...any) ([][]string, error)
	Execute(unfmt string, arg ...any) (int64, error)
	Begin() error
	Commit() error
	Rollback() error
	ReleaseCon()
}

func dbSelect(query string, db *sql.DB) ([][]string, error) {
	row_values := make([][]string, 0)

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
		col_values := make([]any, len(columns))
		string_values := make([]string, len(columns))

		for i := range col_values {
			col_values[i] = new(sql.NullString)
		}

		err := rows.Scan(col_values...)
		if err != nil {
			return nil, err
		}

		for i, value := range col_values {
			string_values[i] = value.(*sql.NullString).String
		}

		row_values = append(row_values, string_values)
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

func getCon(database Database) error {
	switch database := database.(type) {
	case *MysqlDatabase:
		database.mutex.Lock()
		defer database.mutex.Unlock()

		if database.maxConnections == 0 {
			return nil
		}

		if database.maxConnections <= database.numConnections {
			return Err_CON_NOT_AVALIABLE
		}
		database.numConnections++
	}
	return nil
}

func releaseCon(database Database) {
	switch database := database.(type) {
	case *MysqlDatabase:
		database.mutex.Lock()
		defer database.mutex.Unlock()

		if database.numConnections > 0 {
			database.numConnections--
		}
	}
}

func close(database Database) error {
	switch database := database.(type) {
	case *MysqlDatabase:
		return database.db.Close()
	}
	return nil
}
