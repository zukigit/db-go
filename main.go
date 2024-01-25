package main

import (
	"fmt"

	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")

	db.DBconnect()
	db.DBbegin()
	afftected_rows, _ := db.DBexec("INSERT INTO album (title, artist, price) VALUES ('%s', '%s', %d)", "misaki", "zuki", 69)
	fmt.Println("afftected rows:", afftected_rows)
	db.DBrollback()

	result, err := db.DBselect("select * from album")
	if err == nil {
		fmt.Println("result:", result)
	}
}
