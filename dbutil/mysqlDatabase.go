package dbutil

import (
	"database/sql"
	"fmt"
)

type MysqlDatabase struct {
	db             *sql.DB
	isInTranx      *bool
	dataSourceName string
	err            error
}

func NewMysqlDatabase(dataSourceName string) *MysqlDatabase {
	notInTranx := false

	return &MysqlDatabase{dataSourceName: dataSourceName, isInTranx: &notInTranx}
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
	return mysql.db.Close()
}

func (mysql *MysqlDatabase) Select(unfmt string, arg ...any) ([][]string, error) {
	if mysql.db == nil {
		return nil, Err_DB_NOT_INIT
	}
	query := fmt.Sprintf(unfmt, arg...)
	return dbSelect(query, mysql.db)
}

func (mysql *MysqlDatabase) Execute(unfmt string, arg ...any) (int64, error) {
	if mysql.db == nil {
		return 0, Err_DB_NOT_INIT
	}
	query := fmt.Sprintf(unfmt, arg...)
	return dbExecute(query, mysql.db)
}

func (mysql *MysqlDatabase) Begin() error {
	if *mysql.isInTranx {
		return Err_DB_MULTIPLE_TRANSACTIONS
	}
	if _, err := Execute("START TRANSACTION;"); err != nil {
		return err
	}

	*mysql.isInTranx = true
	return nil
}

func (mysql *MysqlDatabase) Commit() error {
	if !*mysql.isInTranx {
		return Err_DB_NO_TRANSACTION
	}
	if _, err := Execute("COMMIT;"); err != nil {
		return err
	}

	*mysql.isInTranx = false
	return nil
}

func (mysql *MysqlDatabase) Rollback() error {
	if !*mysql.isInTranx {
		return Err_DB_NO_TRANSACTION
	}
	if _, err := Execute("ROLLBACK;"); err != nil {
		return err
	}

	*mysql.isInTranx = false
	return nil
}
