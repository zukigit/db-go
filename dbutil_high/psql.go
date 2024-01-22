package dbutil_high

import (
	"github.com/zukigit/db-go/dbutil"
)

type Psql struct {
	DButil *dbutil.DButil
}

func GetInstance_PSQL() *Psql {
	return &Psql{}
}

func (psql *Psql) GetUtilInstance(DBusername string, DBpasswd string, DBname string, DBhost string, DBport string, DBtype string) {
	psql.DButil = dbutil.GetInstance(DBusername, DBpasswd, DBname, DBhost, DBport, DBtype)
}

func (psql *Psql) DBconnect() error {
	return nil
}

func (psql *Psql) DBselect(unfmt string, arg ...any) ([][]interface{}, error) {
	return nil, nil
}
