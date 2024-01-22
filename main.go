package main

import "github.com/zukigit/db-go/dbutil_high"

func main() {
	db := dbutil_high.GetInstance_MYSQL("root", "zabbix", "golearn", "", "")
	db.DBconnect()
	db.DBselect("select * from album")
}
