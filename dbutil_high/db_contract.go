package dbutil_high

type DBcontract interface {
	GetUtilInstance()
	DBconnect() error
	DBselect(unfmt string, arg ...any) ([][]interface{}, error)
}
