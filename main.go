package main

import (
	"fmt"

	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")
	db.DBconnect()

	result, err := db.DBselect("select * from album a where a.title = '%s'", "Blue Train")
	if err == nil {
		fmt.Println("result:", result)
	}
	afftected_rows, _ := db.DBexec("delete from album where title = '%s'", "zuki")
	fmt.Println("afftected rows:", afftected_rows)
}
