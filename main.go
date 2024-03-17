package main

import (
	"fmt"

	db "github.com/zukigit/db-go/dbutil"
)

func main() {
	DBHOST := "moon"
	DBUSER := "root"
	DBPASSWORD := "zabbixx"
	DBNAME := "test"
	DBPORT := 0
	DBCONTIMEOUT := 1

	err := db.Connect_mysql(DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, DBCONTIMEOUT)
	if err != nil {
		fmt.Printf("We get error in connecting Database. Err: %s\n", err.Error())
	} else {
		fmt.Println("Connection Successful!")
	}
}
