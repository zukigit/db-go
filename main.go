package main

import (
	"fmt"

	"github.com/zukigit/db-go/dbutil"
)

func main() {
	var err error
	DBHOST := ""
	DBUSER := "root"
	DBPASSWORD := "zabbix"
	DBNAME := "zabbix"
	DBPORT := 3306

	// need to call only once unless you want to change database config
	dbutil.Init_mysql(
		DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT)

	db, err := dbutil.Connect()
	if err != nil {
		fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
	}

	effected_rows, err := db.Execute("insert into hosts (hostid, description) values(%d, '%s');",
		7071, "")
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		fmt.Println("effected_rows:", effected_rows)
	}

	rows, err := db.Select("select hostid from hosts where hostid = %d", 7071)
	if err != nil {
		fmt.Printf("Error in select. Err: %s\n", err.Error())
	}
	for _, row := range rows {
		fmt.Println(row)
	}

	if err = db.Close(); err != nil {
		fmt.Printf("Error in closing Database. (%s)\n", err.Error())
	}

	// //start transaction
	// if err = db.Begin(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

	// //db execute
	// effected_rows, err := db.Execute("insert into hosts (hostid, description) values(%d, '%s');",
	// 	7071, "")
	// if err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// } else {
	// 	fmt.Println("effected_rows:", effected_rows)
	// }

	// //db commit
	// if err = db.Commit(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

	// //db rollback
	// if err = db.Rollback(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// }

}
