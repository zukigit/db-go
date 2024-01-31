package dbutil_high

import (
	"fmt"
	"sync"

	"github.com/zukigit/db-go/dbutil"
)

type Database struct {
	connection   *dbutil.Connection
	DBisInTx bool
}
var once sync.Once


func getDataSource(dbUsername string, dbPasswd string, dbName string, dbHost string, dbPort string) string{
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUsername, dbPasswd, dbHost, dbPort, dbName)
	
	return dataSourceName
}

func dbInit() *Database {
	var instance *Database

	once.Do(
		func() {
			instance = new(Database)
		},
	)
	instance.DBisInTx = false
	return instance
}

// DBinit_MYSQL returns mysql Database pointer
//
// Only the first three params are mandatory. You can leave the rest as empty string for default values.
func DBinit_MYSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) (error, *Database) {
	DBtype := "mysql"
	if DBport == "" {
		DBport = "3306"
	}
	if DBhost == "" {
		DBhost = "localhost"
	}

	err := dbutil.SetDBsource(getDataSource(DBusername, DBpasswd, DBname, DBhost, DBport), DBtype)
	if err != nil {
		return err, nil
	}

	return err, dbInit()
}

// DBinit_MYSQL returns psql Database pointer
//
// Only the first three params are mandatory. You can leave the rest as empty string for default values.
func DBinit_PSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) (error, *Database) {
	DBtype := "postgres"
	if DBport == "" {
		DBport = "5432"
	}
	if DBhost == "" {
		DBhost = "localhost"
	}

	err := dbutil.SetDBsource(getDataSource(DBusername, DBpasswd, DBname, DBhost, DBport), DBtype)
	if err != nil {
		return err, nil
	}

	return err, dbInit()
}

// func (database *Database) DBconnect() error {
// 	return database.DButil.DBconnect()
// }

// func (database *Database) DBclose() error {
// 	return database.DButil.DBclose()
// }

// func (database *Database) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
// 	query := fmt.Sprintf(unfmt, arg...)
// 	return database.DButil.DBselect(query)
// }

// func (database *Database) DBexec(unfmt string, arg ...any) (int64, error) {
// 	query := fmt.Sprintf(unfmt, arg...)
// 	return database.DButil.DBexec(query)
// }

// func (database *Database) DBbegin() error {
// 	database.DButil.DBbegin()

// 	database = &Database{
// 		DBisInTx: true,
// 		DButil: &dbutil.DButil{
// 			Tx: database.DButil.Tx,
// 		},
// 	}
// 	return nil
// }

// func (database *Database) DBcommit() error {

// 	return database.DButil.DBcommit()
// }

// func (database *Database) DBrollback() error {
// 	database.DBexec("rollback;")
// 	return nil
// }
