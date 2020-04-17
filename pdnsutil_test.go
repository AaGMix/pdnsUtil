package pdnsutil

import (
	"fmt"
	"github.com/AaGMix/pdnsUtil/database/mysqlPack"
	"testing"
)

var C, _ = New(mysqlPack.Config{
	UserName: "root",
	Password: "b7Cht4zJ9",
	Host:     "127.0.0.1",
	Port:     "3366",
	DBName:   "pdns",
	Options:  "charset=utf8",
})

func TestZonesUpdate(t *testing.T) {
	id, err := C.Zones().Update("192.168.120.3", "k3.router.lan")
	if err != nil {
		t.Errorf("Update() Fail: %+v ", err)
		return
	}
	fmt.Println(id)
	return
}
