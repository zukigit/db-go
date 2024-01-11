package main

import (
	"fmt"

	db_util "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := db_util.GetDBsource_MYSQL("root", "zabbix", "golearn", "", "")
	db.DBconnect()
	rows, _ := db.DBselect("select * from album a where a.title = '%s'", "Blue Train")

	for _, value := range rows {
		fmt.Println(value)
	}
}
