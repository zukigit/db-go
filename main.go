package main

import (
	"fmt"

	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")
	db.DBconnect()

	result, err := db.DBselect("select * from album")
	if err == nil {
		fmt.Println("result:", result)
	}
}
