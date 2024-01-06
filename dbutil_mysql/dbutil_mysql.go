package dbutilmysql

import (
	"fmt"
	"zuki/main/dbutil"
)

func chckDBsource(dbsource dbutil.DBsource) dbutil.DBsource {
	dbsource.DBtype = "mysql"
	//check for port and db host
	if dbsource.DBport == "" {
		dbsource.DBport = "3306"
	}
	if dbsource.DBhost == "" {
		dbsource.DBhost = "localhost"
	}

	return dbsource
}

func DBconnect(dbsource dbutil.DBsource) {
	err := dbutil.DBconnect(chckDBsource(dbsource))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}
}
