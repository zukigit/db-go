package dbutil_high

type DBcontract interface {
	GetInstance()
	DBconnect() error
	DBselect(unfmt string, arg ...any) ([][]interface{}, error)
}
