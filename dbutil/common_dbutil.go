package dbutil

import "errors"

var Err_DB_NOT_CONNECTED = errors.New("database is not connected, use GetCon() first")
var Err_DB_NOT_INIT = errors.New("database is not initialized, use Init_() first")
var Err_DB_MULTIPLE_INIT = errors.New("multiple intiliaztion detected, use ReleaseCon() first")
var Err_DB_MULTIPLE_TRANSACTIONS = errors.New("doesn't not support multile transaction, you have to close one transaction using Close()")
var Err_DB_NO_TRANSACTION = errors.New("no transaction is detected, use begin() to start the transaction")

var Err_CON_NOT_AVALIABLE = errors.New("connection is not avaliable yet")
