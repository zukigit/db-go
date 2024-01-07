package dbutilmysql

import (
	"fmt"

	"github.com/zukigit/db-go/dbutil"
)

func GetDBsource(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) dbutil.DBsource {
	return dbutil.ChckDBsource(dbutil.GetDBsource(DBusername, DBpasswd, DBname, DBhost, DBport), dbutil.MYSQL)
}

func DBconnect(dbsource dbutil.DBsource) {
	err := dbutil.DBconnect(dbsource)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}
}
