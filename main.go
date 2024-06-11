package main

import (
	"database/sql"
	"fmt"
	"reflect"

	db "github.com/zukigit/db-go/dbutil"
)

func main() {
	var err error

	DBHOST := ""
	DBUSER := "root"
	DBPASSWORD := "zabbix"
	DBNAME := "zabbix"
	DBPORT := 0
	DBCONTIMEOUT := 1

	err = db.Connect_mysql(
		DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, DBCONTIMEOUT)
	if err != nil {
		fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
	}

	result, err := db.Select("select  host from hosts h where hostid = %d;", 10050)
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		fmt.Println("result:", result)
		for _, value := range result {
			fmt.Println("values:", reflect.TypeOf(value[0]))
			fmt.Println("values:", value[0].(sql.NullString).String)
		}
	}

	err = db.Close()
	if err != nil {
		fmt.Printf("Error in closing Database. (%s)\n", err.Error())
	}

	//db select
	// result, err := db.Select("select title from album where id = '%d';", 1)
	// if err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// } else {
	// 	for _, values := range result {
	// 		fmt.Println("values:", values[0])
	// 	}
	// }

	//db begin
	// if err = db.Begin(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

	//db execute
	// effected_rows, err := db.Execute("INSERT  INTO album (id, title, artist, price) VALUES (%d, '%s', '%s', %d);",
	// 	43, ":)))))", "me", 69)
	// if err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// } else {
	// 	fmt.Println("effected_rows:", effected_rows)
	// }

	// //db rollback
	// if err = db.Rollback(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

	//db commit
	// if err = db.Commit(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }
}
