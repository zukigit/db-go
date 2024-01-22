package main

import "github.com/zukigit/db-go/dbutil_high"

func main() {
	db := dbutil_high.GetInstance_MYSQL()
	db.GetUtilInstance("root", "zabbix", "golearn", "", "")
}
