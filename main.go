package main

import (
	mysql "github.com/zukigit/db-go/dbutil_mysql"
)

func main() {
	mysql.DBconnect(mysql.GetDBsource("root", "zabbix", "golearn", "", "")) //only the first three params are mandatory field. you can leave blank for optinal fields.
}
