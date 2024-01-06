package dbutil

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBsource struct {
	DBusername string
	DBpasswd   string
	DBname     string
	DBhost     string
	DBtype     string
	DBport     string
}

func DBconnect(dbsource DBsource) error {
	var db *sql.DB
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbsource.DBusername, dbsource.DBpasswd, dbsource.DBhost, dbsource.DBport, dbsource.DBname)

	db, err = sql.Open(dbsource.DBtype, dataSourceName)
	if err != nil {
		return errors.New(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		db.Close()
		return errors.New(pingErr.Error())
	}

	return nil
}
