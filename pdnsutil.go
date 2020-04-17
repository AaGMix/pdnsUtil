package pdnsutil

import (
	"github.com/AaGMix/pdnsUtil/database/mysqlPack"
	"github.com/AaGMix/pdnsUtil/zones"
)

type control struct {
	mysqlConfig mysqlPack.Config
	zones       zones.Control
}

func New(config mysqlPack.Config) (Control, error) {
	c := control{
		mysqlConfig: config,
	}
	c.zones, _ = zones.New(c.mysqlConfig)
	return &c, nil
}

func (c *control) Zones() zones.Control {
	return c.zones
}
