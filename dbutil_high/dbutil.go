package dbutil_high

import "github.com/zukigit/db-go/dbutil"

// GetDBsource creates instance MYSQL DBsource object.
//
// Only the first three params are mandatory fields and the others are optional. You can leave empty string for optional fields. So the empty fields will be set with default values.
func GetDBsource_MYSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) dbutil.DBsource {
	return dbutil.ChckDBsource(dbutil.GetDBsource(DBusername, DBpasswd, DBname, DBhost, DBport, dbutil.MYSQL))
}

// GetDBsource creates instance POSTGRESQL DBsource object.
//
// Only the first three params are mandatory fields and the others are optional. You can leave empty string for optional fields. So the empty fields will be set with default values.
func GetDBsource_PSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) dbutil.DBsource {
	return dbutil.ChckDBsource(dbutil.GetDBsource(DBusername, DBpasswd, DBname, DBhost, DBport, dbutil.POSTGRESQL))
}
