package dbutilmysql

import (
	"github.com/zukigit/db-go/dbutil"
)

// GetDBsource creates instance DBsource object.
//
// Only the first three params are mandatory fields and the others are optional. You can leave empty string for optional fields. So the empty fields will be set with default values.
func GetDBsource(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) dbutil.DBsource {
	return dbutil.ChckDBsource(dbutil.GetDBsource(DBusername, DBpasswd, DBname, DBhost, DBport), dbutil.MYSQL)
}

func DBconnect(dbsource dbutil.DBsource) error{
	return dbutil.DBconnect(dbsource)
}
