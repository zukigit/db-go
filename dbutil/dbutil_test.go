package dbutil

import (
	"testing"
)

func TestDButil(t *testing.T) {
	Init_mysql("", "", "", "", 0, 0)
	if _, err := GetCon(); err == nil {
		t.Errorf("dbutil.GetCon() should get failed")
	}

	Init_mysql_DSN("sdfsdfnkhjiu", 0)
	if _, err := GetCon(); err == nil {
		t.Errorf("dbutil.GetCon() should get failed")
	}

	Init_mysql("", "zabbix", "zabbix", "zabbix", 0, 0)
	if _, err := GetCon(); err != nil {
		t.Errorf("dbutil.GetCon() get failed, err: %s", err.Error())
	}

	Init_mysql_DSN("zabbix:zabbix@tcp/zabbix?allowNativePasswords=false&checkConnLiveness=false&maxAllowedPacket=0", 0)
	if _, err := GetCon(); err != nil {
		t.Errorf("dbutil.GetCon() get failed, err: %s", err.Error())
	}
}
