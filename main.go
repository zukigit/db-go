package main

import (
	"fmt"
	"zuki/main/dbutil"
	mysql "zuki/main/dbutil_mysql"
)

func main() {
	fmt.Println("This is main")

	dbsource := dbutil.DBsource{
		DBusername: "root",
		DBpasswd:   "zabbix",
		DBname:     "golearn",
		DBhost:     "moon",
	}
	mysql.DBconnect(dbsource)
}
