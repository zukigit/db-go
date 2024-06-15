package dbutil

import "errors"

var Err_DB_NOT_INIT = errors.New("database is not initialized, use Connect_() first")
var Err_DB_MULTIPLE_INIT = errors.New("multiple intiliaztion detected, use Close() first")
var Err_DB_MULTIPLE_TRANSACTIONS = errors.New("doesn't not support multile transaction, you have to close one transaction using Close()")
var Err_DB_NO_TRANSACTION = errors.New("no transaction is detected, use begin() to start the transaction")
