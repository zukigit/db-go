package dbutil_high

import (
	"fmt"

	"github.com/zukigit/db-go/dbutil"
)

type Mysql struct {
	DButil *dbutil.DButil
}

func GetInstance_MYSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) *Mysql {
	DBtype := "mysql"
	return &Mysql{
		DButil: dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype),
	}
}

func (mysql *Mysql) DBconnect() error {
	return mysql.DButil.DBconnect()
}

func (mysql *Mysql) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	query := fmt.Sprintf(unfmt, arg...)
	return mysql.DButil.DBselect(query)
}
