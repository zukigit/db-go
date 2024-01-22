package dbutil_high

import (
	"fmt"
	"sync"

	"github.com/zukigit/db-go/dbutil"
)

type Mysql struct {
	DButil *dbutil.DButil
}

func DBinit_MYSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) *Mysql {
	var once sync.Once
	var instance *Mysql
	DBtype := "mysql"

	if DBhost == "" {
		DBhost = "localhost"
	}
	if DBport == "" {
		DBport = "3306"
	}

	once.Do(
		func() {
			instance = new(Mysql)
		},
	)
	instance.DButil = dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)

	return instance
}

func (mysql *Mysql) DBconnect() error {
	return mysql.DButil.DBconnect()
}

func (mysql *Mysql) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	query := fmt.Sprintf(unfmt, arg...)
	return mysql.DButil.DBselect(query)
}
