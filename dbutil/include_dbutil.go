package dbutil

import "errors"

const (
	//db types
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"

	//data types
	VARCHAR  = "VARCHAR"
	TEXT     = "TEXT"
	NVARCHAR = "NVARCHAR"
	DECIMAL  = "DECIMAL"
	BOOL     = "BOOL"
	INT      = "INT"
	BIGINT   = "BIGINT"
)

var Err_DB_NOT_INIT = errors.New("Err_DB_NOT_INIT")
var Err_DB_MULTIPLE_TRANSACTIONS = errors.New("Err_DB_MULTIPLE_TRANSACTIONS")
var Err_DB_NO_TRANSACTION = errors.New("Err_DB_NO_TRANSACTION")