package mysqlPack

type Control interface {
	Ping() error

	Insert(sqlStr string, args ...interface{}) (int64, error)
	UpdateOrDelete(sqlStr string, args ...interface{}) (int64, error)
	Query(sqlStr string, args ...interface{}) (*[]map[string]string, error)
	QueryRow(sqlStr string, args ...interface{}) (*map[string]string, error)

	DBClose() error
}
