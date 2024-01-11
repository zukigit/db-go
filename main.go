package main

import (
	"fmt"

	mysql "github.com/zukigit/db-go/dbutil_mysql"
)

func main() {
	mysql.DBconnect(mysql.GetDBsource("root", "zabbix", "golearn", "", ""))
	rows, _ := mysql.DBselect("select * from album a where a.title = '%s'", "Blue Train")

	for _, value := range rows {
		fmt.Println(value)
	}
}
