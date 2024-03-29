package main

import (
	"fmt"

	db "github.com/zukigit/db-go/dbutil"
)

func main() {
	var err error

	DBHOST := "10.1.9.30"
	DBUSER := "zabbix"
	DBPASSWORD := "zabbix"
	DBNAME := "zabbix"
	DBPORT := 0
	DBCONTIMEOUT := 1

	//db connect
	err = db.Connect_mysql(
		DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, DBCONTIMEOUT)
	if err != nil {
		fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
	}

	//db select
	result, err := db.Select("select hostid from hosts where host = '%s';", "rhel6")
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		for _, values := range result {
			fmt.Println("values:", values[0])
		}
	}

	//db begin
	if err = db.Begin(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	}

	//db execute
	effected_rows, err := db.Execute("INSERT  INTO test (id, title, artist, price) VALUES (%d, '%s', '%s', %d);",
		43, ":)))))", "me", 69)
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		fmt.Println("effected_rows:", effected_rows)
	}

	// //db rollback
	// if err = db.Rollback(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

	//db commit
	if err = db.Commit(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	}
}
