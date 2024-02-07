package main

import (
	"fmt"

	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")
	// db.DBexec("INSERT INTO album (title, artist, price) VALUES ('%s', '%s', %d)", "zuki", "zuki", 69)
	db.DBexec("delete from album where id = '%d'", 4)
	result, _ := db.DBselect("select * from album a")
	fmt.Println("result :", result)
	db.DBclose()
}
