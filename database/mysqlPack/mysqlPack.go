package mysqlPack

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Config struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
	Options  string
}

type control struct {
	config   Config
	database *sql.DB
}

func New(cfg Config) (Control, error) {
	c := control{
		config: cfg,
	}
	var err error
	path := strings.Join([]string{c.config.UserName, ":", c.config.Password,
		"@tcp(", c.config.Host, ":", c.config.Port, ")/", c.config.DBName, "?", c.config.Options}, "")

	c.database, err = sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}
	// 设置数据库最大连接数
	c.database.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	c.database.SetMaxIdleConns(10)

	// 验证连接
	if err := c.database.Ping(); err != nil {
		fmt.Println("opon database fail")
		return nil, err
	}

	return &c, nil
}

func (c *control) Ping() error {
	if err := c.database.Ping(); err != nil {
		return err
	}
	return nil
}

func (c *control) Insert(sqlStr string, args ...interface{}) (int64, error) {

	//开启事务
	tx, err := c.database.Begin()
	if err != nil {
		return -1, fmt.Errorf("tx fail: %s", err)
	}
	//准备sql语句
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return -1, fmt.Errorf("prepare fail: %s", err)
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("exec fail: %s", err)
	}
	//将事务提交
	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("Commit Transaction Fail, Error: %+v ", err)
	}
	//获得上一个插入自增的id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("get last insert id fail: %s", err)
	}
	return id, nil
}
func (c *control) UpdateOrDelete(sqlStr string, args ...interface{}) (int64, error) {
	//开启事务
	tx, err := c.database.Begin()
	if err != nil {
		return -1, fmt.Errorf("tx fail: %s", err)
	}
	//准备sql语句
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return -1, fmt.Errorf("prepare fail: %s", err)
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("exec fail: %s", err)
	}
	//将事务提交
	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("Commit Transaction Fail, Error: %+v ", err)
	}
	//获得上一个插入自增的id
	id, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("get last insert id fail: %s", err)
	}
	return id, nil
}
func (c *control) Query(sqlStr string, args ...interface{}) (*[]map[string]string, error) {
	//执行查询语句
	rows, err := c.database.Query(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("query fail: %s", err)
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("columns fail: %s", err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//循环读取结果
	for rows.Next() {
		//将每一行的结果都赋值到一个user对象中
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("rows fail: %s", err)
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil
}

func (c *control) QueryRow(sqlStr string, args ...interface{}) (*map[string]string, error) {
	//执行查询语句
	rows, err := c.database.Query(sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("query fail: %s", err)
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("columns fail: %s", err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make(map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//循环读取结果
	for rows.Next() {
		//将每一行的结果都赋值到一个user对象中
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("rows fail: %s", err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break
	}
	return &ret, nil
}

func (c *control) DBClose() error {
	return c.database.Close()
}
