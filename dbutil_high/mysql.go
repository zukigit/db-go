package dbutil_high

import (
	"github.com/zukigit/db-go/dbutil"
)

type Mysql struct {
	DButil *dbutil.DButil
}

func (mysql *Mysql) GetInstance(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string, DBtype string) {
	mysql.DButil = dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)
}

func (mysql *Mysql) DBconnect() error {
	return nil
}

func (mysql *Mysql) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	return nil, nil
}
