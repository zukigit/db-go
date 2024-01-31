package main

import (
	dbutil "github.com/zukigit/db-go/dbutil_high"
)

func main() {
	dbutil.DBinit_MYSQL("root", "zabbix", "golearn", "", "")

	// db.DBconnect()
	// db.DBbegin()

	// tx_afftected_rows, _ := db1.DBexec("INSERT INTO album (title, artist, price) VALUES ('%s', '%s', %d)", "from trax", "zuki", 69)
	// fmt.Println("afftected rows:", tx_afftected_rows)

	// afftected_rows, _ := db.DBexec("INSERT INTO album (title, artist, price) VALUES ('%s', '%s', %d)", "normal", "zuki", 69)
	// fmt.Println("afftected rows:", afftected_rows)

	// db.DBrollback()

	// result, err := db.DBselect("select * from album")
	// if err == nil {
	// 	fmt.Println("result:", result)
	// }
}
