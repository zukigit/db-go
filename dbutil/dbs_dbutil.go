package dbutil

import (
	"database/sql"
	"fmt"
)

type MysqlDatabase struct {
	db        *sql.DB
	err       *error
	isInTranx *bool
}

type PsqlDatabase struct {
}

func (db MysqlDatabase) Select(unfmt string, arg ...any) ([][]interface{}, error) {
	if isDBinit() {
		query := fmt.Sprintf(unfmt, arg...)
		return dbSelect(query)
	}
	return nil, Err_DB_NOT_INIT
}

func (db MysqlDatabase) Begin() error {
	if !*db.isInTranx {
		if *db.err = dbBegin("START TRANSACTION; "); db.err != nil {
			*db.isInTranx = true
			return nil
		} else {
			return *db.err
		}
	}
	return Err_DB_MULTIPLE_TRANSACTIONS
}

func (mysql MysqlDatabase) Close() error {
	return mysql.db.Close()
}

func (db MysqlDatabase) Commit() error {
	if *db.isInTranx {
		return dbCommit()
	} else {
		return Err_DB_NO_TRANSACTION
	}
}

func (db MysqlDatabase) Execute(unfmt string, arg ...any) (int64, error) {
	if isDBinit() {
		query := fmt.Sprintf(unfmt, arg...)
		return dbExecute(query)
	}
	return 0, Err_DB_NOT_INIT
}

func (db MysqlDatabase) Rollback() error {
	if *db.isInTranx {
		return dbRollback()
	} else {
		return Err_DB_NO_TRANSACTION
	}
}
