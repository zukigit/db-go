package dbutil

import "errors"

const (
	//db types
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"

	//data types
	VARCHAR            = "VARCHAR"
	TEXT               = "TEXT"
	NVARCHAR           = "NVARCHAR"
	DECIMAL            = "DECIMAL"
	BOOL               = "BOOL"
	INT                = "INT"
	UNSIGNED_INT       = "UNSIGNED INT"
	TINYINT            = "TINYINT"
	UNSIGNED_TINYINT   = "UNSIGNED TINYINT"
	SMALLINT           = "SMALLINT"
	UNSIGNED_SMALLINT  = "UNSIGNED SMALLINT"
	MEDIUMINT          = "MEDIUMINT"
	UNSIGNED_MEDIUMINT = "UNSIGNED MEDIUMINT"
	BIGINT             = "BIGINT"
	UNSIGNED_BIGINT    = "UNSIGNED BIGINT"
)

var Err_DB_NOT_INIT = errors.New("Err_DB_NOT_INIT")
var Err_DB_MULTIPLE_TRANSACTIONS = errors.New("Err_DB_MULTIPLE_TRANSACTIONS")
var Err_DB_NO_TRANSACTION = errors.New("Err_DB_NO_TRANSACTION")
var Err_UNDEFINED_COLLUMN_TYPE = errors.New("Err_UNDEFINED_COLLUMN_TYPE")
