package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	db "github.com/zukigit/db-go/dbutil"
)

func main() {
	var err error

	// DBHOST := ""
	// DBUSER := "root"
	// DBPASSWORD := "zabbix"
	// DBNAME := "zabbix"
	// DBPORT := 3306

	// err = db.Connect_mysql(
	// 	DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT)
	// if err != nil {
	// 	fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
	// }

	cfg := mysql.Config{
		User:   "zabbix",
		Passwd: "zabbix",
		Net:    "tcp",
		Addr:   "",
		DBName: "zabbix",
	}
	err = db.Connect_mysql(cfg.FormatDSN())
	if err != nil {
		fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
	}

	// //db select
	result, err := db.Select("select  host from hosts h where hostid = %d;", 10050)
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		fmt.Println("result:", result)
	}

	// //start transaction
	if err = db.Begin(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	}

	// //db execute
	effected_rows, err := db.Execute("insert into hosts (hostid, description) values(%d, '%s');",
		7071, "")
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		fmt.Println("effected_rows:", effected_rows)
	}

	// //db commit
	if err = db.Commit(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	}

	// //db rollback
	if err = db.Rollback(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	}

	// //close database connection
	if err = db.Close(); err != nil {
		fmt.Printf("Error in closing Database. (%s)\n", err.Error())
	}
}
