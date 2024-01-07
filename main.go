package main

import (
	"fmt"

	"github.com/zukigit/db-go/dbutil"
	mysql "github.com/zukigit/db-go/dbutil_mysql"
)

func main() {
	fmt.Println("This is main")

	dbsource := dbutil.DBsource{
		DBusername: "root",
		DBpasswd:   "zabbix",
		DBname:     "golearn",
	}
	mysql.DBconnect(dbsource)
}
