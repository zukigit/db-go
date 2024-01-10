package main

import (
	mysql "github.com/zukigit/db-go/dbutil_mysql"
)

func main() {
	mysql.DBconnect(mysql.GetDBsource("zabbix", "zabbix", "zabbix", "192.168.238.128", ""))
	mysql.DBselect("name: %s", "zuki")
}
