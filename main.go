package main

import (
	"fmt"
	"time"

	"github.com/zukigit/db-go/dbutil"
)

func doTest() {
	var err error

	db, err := dbutil.GetConnection()
	if err != nil {
		fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
		return
	}

	if err = db.Begin(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
		return
	}

	effected_rows, err := db.Execute("insert into hosts (hostid, description) values(%d, '%s');",
		7075, "")
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
		return
	} else {
		fmt.Println("effected_rows:", effected_rows)
	}

	// if err = db.Commit(); err != nil {
	// 	fmt.Printf("Query get failed, error: %s\n", err.Error())
	// 	return
	// }

	if err = db.Rollback(); err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
		return
	}

	rows, err := db.Select("select hostid from hosts where hostid = %d", 7073)
	if err != nil {
		fmt.Printf("Error in select. Err: %s\n", err.Error())
		return
	}
	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println("task succeeded!!!!")
	db.ReleaseCon()
}

func main() {

	DBHOST := ""
	DBUSER := "root"
	DBPASSWORD := "zabbix"
	DBNAME := "zabbix"
	DBPORT := 3306

	dbutil.Init_mysql(
		DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, 1)

	go doTest()
	time.Sleep(1 * time.Second)
	go doTest()
	time.Sleep(1 * time.Second)
	go doTest()
	time.Sleep(1 * time.Second)

	dbutil.Close()
}
