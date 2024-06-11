package dbutil

import "errors"

const (
	//db types
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"

	//column types
	VARCHAR            = "VARCHAR"
	NVARCHAR           = "NVARCHAR"
	TEXT               = "TEXT"
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
	DECIMAL            = "DECIMAL"
	BOOL               = "BOOL"
	FLOAT              = "FLOAT"
	DOUBLE             = "DOUBLE"
	DATE               = "DATE"
	DATETIME           = "DATETIME"
	TIMESTAMP          = "TIMESTAMP"
	TIME               = "TIME"
	YEAR               = "YEAR"
	CHAR               = "CHAR"
	BINARY             = "BINARY"
	VARBINARY          = "VARBINARY"
	BLOB               = "BLOB"
	MEDIUMBLOB         = "MEDIUMBLOB"
	LONGBLOB           = "LONGBLOB"
	JSON               = "JSON"
	ENUM               = "ENUM"
	SET                = "SET"
	POINT              = "POINT"
	GEOMETRY           = "GEOMETRY"
)

var Err_DB_NOT_INIT = errors.New("Err_DB_NOT_INIT")
var Err_DB_MULTIPLE_TRANSACTIONS = errors.New("Err_DB_MULTIPLE_TRANSACTIONS")
var Err_DB_NO_TRANSACTION = errors.New("Err_DB_NO_TRANSACTION")
var Err_UNDEFINED_COLLUMN_TYPE = errors.New("Err_UNDEFINED_COLLUMN_TYPE")
