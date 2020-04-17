package zones

import (
	"fmt"
	"github.com/AaGMix/pdnsUtil/database/mysqlPack"
)

type control struct {
	mysqlP mysqlPack.Control
}

func New(cfg mysqlPack.Config) (Control, error) {
	mysqlControl, err := mysqlPack.New(cfg)
	if err != nil {
		return nil, err
	}
	c := control{
		mysqlP: mysqlControl,
	}
	if err = c.mysqlP.Ping(); err != nil {
		return nil, fmt.Errorf("Ping() Fail, Error : %+v ", err)
	}
	return &c, nil
}

func (c *control) CClose() error {
	return c.mysqlP.DBClose()
}
