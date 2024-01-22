package dbutil_high

import (
	"github.com/zukigit/db-go/dbutil"
)

type Psql struct {
	DButil *dbutil.DButil
}

func GetInstance_PSQL(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string) *Psql {
	DBtype := "mysql"
	return &Psql{
		DButil: dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype),
	}
}

func (psql *Psql) DBconnect() error {
	return nil
}

func (psql *Psql) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	return nil, nil
}
