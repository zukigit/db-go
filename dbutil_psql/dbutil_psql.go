package dbutilpsql

import (
	"database/sql"
	"fmt"

	"github.com/zukigit/db-go/dbutil"
)

var db *sql.DB

func DBconnect(dbsource dbutil.DBsource) {
	fmt.Println("It's db connect")

	dbsource.DBtype = "postgres"
	err := dbutil.DBconnect(dbsource, db)

	if err != nil {
		fmt.Println(err)
	}
}
