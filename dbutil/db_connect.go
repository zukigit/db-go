package dbutil

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func DBconnect() {
	fmt.Println("It's db connect")

	db, err = sql.Open("mysql", "root:zabbix@/golearn")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database is connected!")
}
