package dbutil

import (
	"database/sql"
	"fmt"
)

type MysqlDatabase struct {
	db        *sql.DB
	isInTranx *bool
	dns       string
	// the maximum number of connections in the pool
	// maxConnections int

	// the current number of connections in the pool
	// numConnections int

	// the mutex to synchronize access to the connection pool
	// mutex *sync.Mutex
}

func NewMysqlDatabase(dataSourceName string) (*MysqlDatabase, error) {
	notInTranx := false

	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MysqlDatabase{
		db:        sqlDB,
		isInTranx: &notInTranx,
		dns:       dataSourceName,
	}, nil
}

func (mysql *MysqlDatabase) Ping() error {
	err := mysql.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (mysql *MysqlDatabase) Connect() (database Database, err error) {
	if mysql.db != nil {
		if mysql.db, err = sql.Open("mysql", mysql.dns); err != nil {
			return
		}
	}
	if err = mysql.Ping(); err != nil {
		return
	}
	database = mysql
	return
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
