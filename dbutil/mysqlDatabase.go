package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
)

type MysqlDatabase struct {
	db             *sql.DB
	isInTranx      *bool
	dataSourceName string
	err            error
}

func NewMysqlDatabase(dataSourceName string) *MysqlDatabase {
	return &MysqlDatabase{dataSourceName: dataSourceName}
}

func (mysql MysqlDatabase) Ping() error {
	mysql.err = mysql.db.Ping()
	if mysql.err != nil {
		return mysql.err
	}
	return nil
}

func (mysql *MysqlDatabase) Connect() error {
	if mysql.db, mysql.err = sql.Open("mysql", mysql.dataSourceName); mysql.err != nil {
		return mysql.err
	}

	return mysql.Ping()
}

func (mysql *MysqlDatabase) Close() error {
	if mysql.db == nil {
		fmt.Println("db is nill")
		return errors.New("db is nill")
	}
	return mysql.db.Close()
}

func (mysql *MysqlDatabase) Select(unfmt string, arg ...any) ([][]interface{}, error) {
	if mysql.db != nil { //check whether database is initialized or not
		query := fmt.Sprintf(unfmt, arg...)
		return dbSelect(query, mysql.db)
	}
	return nil, Err_DB_NOT_INIT
}

func (mysql *MysqlDatabase) Execute(unfmt string, arg ...any) (int64, error) {
	if mysql.db != nil {
		query := fmt.Sprintf(unfmt, arg...)
		return dbExecute(query, mysql.db)
	}
	return 0, Err_DB_NOT_INIT
}

func (mysql *MysqlDatabase) Begin() error {
	if !*mysql.isInTranx {
		err := dbBegin("START TRANSACTION;", mysql.db)
		if err == nil {
			*mysql.isInTranx = true
		}
		return err
	}
	return Err_DB_MULTIPLE_TRANSACTIONS
}

func (mysql *MysqlDatabase) Commit() error {
	if !*mysql.isInTranx {
		err := dbCommit("COMMIT;", mysql.db)
		if err == nil {
			*mysql.isInTranx = false
		}
		return err
	}
	return Err_DB_NO_TRANSACTION
}

func (mysql *MysqlDatabase) Rollback() error {
	if !*mysql.isInTranx {
		err := dbRollback("ROLLBACK;", mysql.db)
		if err == nil {
			*mysql.isInTranx = false
		}
		return err
	}
	return Err_DB_NO_TRANSACTION
}
