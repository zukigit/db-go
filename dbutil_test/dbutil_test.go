package dbutiltest

import (
	"testing"

	"github.com/zukigit/db-go/dbutil"
)

func TestDButil(t *testing.T) {
	dbutil.Init_mysql("", "", "", "", 0, 0)
	if _, err := dbutil.GetCon(); err == nil {
		t.Errorf("dbutil.GetCon() should get failed")
	}

	dbutil.Init_mysql_DSN("sdfsdfnkhjiu", 0)
	if _, err := dbutil.GetCon(); err == nil {
		t.Errorf("dbutil.GetCon() should get failed")
	}

	dbutil.Init_mysql("", "zabbix", "zabbix", "zabbix", 0, 0)
	if _, err := dbutil.GetCon(); err != nil {
		t.Errorf("dbutil.GetCon() get failed, err: %s", err.Error())
	}

	dbutil.Init_mysql_DSN("zabbix:zabbix@tcp/zabbix?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", 0)
	if _, err := dbutil.GetCon(); err != nil {
		t.Errorf("dbutil.GetCon() get failed, err: %s", err.Error())
	}
}
