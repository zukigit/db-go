package dbutil

import (
	"database/sql"
	"fmt"
	"sync"
)

type MysqlDatabase struct {
	db        *sql.DB
	isInTranx *bool
	dns       string

	maxConnections int
	numConnections int
	mutex          *sync.Mutex
}

func NewMysqlDatabase(dataSourceName string) *MysqlDatabase {
	notInTranx := false

	return &MysqlDatabase{
		isInTranx: &notInTranx,
		dns:       dataSourceName,
		mutex:     &sync.Mutex{},
	}
}

func (mysql *MysqlDatabase) Ping() error {
	err := mysql.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (mysql *MysqlDatabase) Connect() error {
	var err error

	if err = getCon(mysql); err != nil {
		return err
	}

	mysql.db, err = sql.Open("mysql", mysql.dns)
	if err != nil {
		return err
	}

	return mysql.db.Ping()
}

func (mysql *MysqlDatabase) Close() error {
	if mysql.db == nil {
		return Err_DB_NOT_CONNECTED
	}

	return mysql.db.Close()
}

func (mysql *MysqlDatabase) Select(unfmt string, arg ...any) ([][]string, error) {
	if mysql.db == nil {
		return nil, Err_DB_NOT_CONNECTED
	}
	query := fmt.Sprintf(unfmt, arg...)

	return dbSelect(query, mysql.db)
}

func (mysql *MysqlDatabase) Execute(unfmt string, arg ...any) (int64, error) {
	if mysql.db == nil {
		return 0, Err_DB_NOT_CONNECTED
	}
	query := fmt.Sprintf(unfmt, arg...)

	return dbExecute(query, mysql.db)
}

func (mysql *MysqlDatabase) Begin() error {
	if *mysql.isInTranx {
		return Err_DB_MULTIPLE_TRANSACTIONS
	}
	if _, err := mysql.Execute("START TRANSACTION;"); err != nil {
		return err
	}
	*mysql.isInTranx = true

	return nil
}

func (mysql *MysqlDatabase) Commit() error {
	if !*mysql.isInTranx {
		return Err_DB_NO_TRANSACTION
	}
	if _, err := mysql.Execute("COMMIT;"); err != nil {
		return err
	}
	*mysql.isInTranx = false

	return nil
}

func (mysql *MysqlDatabase) Rollback() error {
	if !*mysql.isInTranx {
		return Err_DB_NO_TRANSACTION
	}
	if _, err := mysql.Execute("ROLLBACK;"); err != nil {
		return err
	}
	*mysql.isInTranx = false

	return nil
}
