package main

import (
	"fmt"

	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	db := dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")

	db.DBconnect()
	fmt.Println("before DB:", db.DButil.DBisInTx)

	db1 := db.DBbegin()

	fmt.Println("DB:", db.DButil.DBisInTx, "DB1:", db1.DButil.DBisInTx)
	// afftected_rows, _ := db.DBexec("INSERT INTO album (title, artist, price) VALUES ('%s', '%s', %d)", "misaki", "zuki", 69)
	// fmt.Println("afftected rows:", afftected_rows)
	// db.DBrollback()

	// result, err := db.DBselect("select * from album")
	// if err == nil {
	// 	fmt.Println("result:", result)
	// }
}
