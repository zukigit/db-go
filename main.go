package main

import (
	mysql "github.com/zukigit/db-go/dbutil_mysql"
)

func main() {
	mysql.DBconnect(mysql.GetDBsource("root", "zabbix", "golearn", "", ""))
	mysql.DBselect("select * from album a")
}
