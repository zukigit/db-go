package dbutilpsql

import (
	"fmt"
	"zuki/main/dbutil"
)

func DBconnect(dbsource dbutil.DBsource) {
	fmt.Println("It's db connect")

	dbsource.DBtype = "postgres"
	err := dbutil.DBconnect(dbsource)

	if err != nil {
		fmt.Println(err)
	}
}
