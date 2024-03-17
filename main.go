package main

import (
	"fmt"

	db "github.com/zukigit/db-go/dbutil"
)

func main() {
	DBHOST := "moon"
	DBUSER := "root"
	DBPASSWORD := "zabbix"
	DBNAME := "test"
	DBPORT := 0
	DBCONTIMEOUT := 1

	err := db.Connect_mysql(DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, DBCONTIMEOUT)
	if err != nil {
		fmt.Printf("We get error in connecting Database. Err: %s\n", err.Error())
	} else {
		fmt.Println("Connection Successful!")
	}

	result, err := db.Select("select * from test;")
	if err != nil {
		fmt.Printf("Query get failed, error: %s\n", err.Error())
	} else {
		for _, values := range result {
			fmt.Println("values:", values[0])
			fmt.Println("values:", values[1])
			fmt.Println("values:", values[2])
			fmt.Println("values:", values[3])
		}
	}
}
