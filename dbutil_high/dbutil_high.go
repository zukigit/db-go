package dbutil_high

import (
	"fmt"
	"sync"

	"github.com/zukigit/db-go/dbutil"
)

type Database struct {
	DButil *dbutil.DButil
}

func DBinit(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string, DBtype string) *Database {
	var once sync.Once
	var instance *Database

	if DBhost == "" {
		DBhost = "localhost"
	}

	once.Do(
		func() {
			instance = new(Database)
		},
	)
	instance.DButil = dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)
	return instance
}

// DBinit_MYSQL returns mysql Database pointer
//
// Only the first three params are mandatory. You can leave the rest as empty string for default values.
func DBinit_MYSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) *Database {
	DBtype := "mysql"
	if DBport == "" {
		DBport = "3306"
	}

	return DBinit(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)
}

// DBinit_MYSQL returns psql Database pointer
//
// Only the first three params are mandatory. You can leave the rest as empty string for default values.
func DBinit_PSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) *Database {
	DBtype := "postgres"
	if DBport == "" {
		DBport = "5432"
	}

	return DBinit(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)
}

func (database *Database) DBconnect() error {
	return database.DButil.DBconnect()
}

func (database *Database) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	query := fmt.Sprintf(unfmt, arg...)
	return database.DButil.DBselect(query)
}

func (database *Database) DBexec(unfmt string, arg ...any) (int64, error) {
	query := fmt.Sprintf(unfmt, arg...)
	return database.DButil.DBexec(query)
}