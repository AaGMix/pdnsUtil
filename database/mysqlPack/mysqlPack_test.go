package mysqlPack

import (
	"fmt"
	"testing"
)

var baseConfig = Config{
	UserName: "root",
	Password: "b7Cht4zJ9",
	Host:     "127.0.0.1",
	Port:     "3366",
	DBName:   "pdns",
	Options:  "charset=utf8",
}

func TestNew(t *testing.T) {

	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name:    "mysqlTest",
			cfg:     baseConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := New(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				defer db.DBClose()
			}
		})
	}
}

func Test_mysqlDB_Ping(t *testing.T) {
	mysqlC, err := New(baseConfig)
	tests := []struct {
		name         string
		mysqlControl Control
		wantErr      bool
	}{
		{
			name:         "Ping test",
			mysqlControl: mysqlC,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = tt.mysqlControl.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mysqlDB_Insert(t *testing.T) {
	mysqlC, err := New(baseConfig)
	if err != nil {
		t.Errorf("New() Run error: %+v\r\n", err)
	}
	var par = []interface{}{3, 1, "k31.router.lan", "A", "192.168.120.2", 10, 0, 0, "", 1}
	var sqlCommand = "INSERT INTO pdns.records VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	id, err := mysqlC.Insert(sqlCommand, par...)
	if err != nil {
		t.Errorf("Insert() Run error: %+v\r\n", err)
	}
	fmt.Printf("Insert data id is: %v\r\n", id)
	defer mysqlC.DBClose()
	return
}

func Test_mysqlDB_Update(t *testing.T) {
	mysqlC, err := New(baseConfig)
	if err != nil {
		t.Errorf("New() Run error: %+v\r\n", err)
	}
	var sqlCommand = "UPDATE pdns.records SET content = ? WHERE name = ?"
	var par = []interface{}{"192.168.120.31", "k31.router.lan"}
	id, err := mysqlC.UpdateOrDelete(sqlCommand, par...)
	if err != nil {
		t.Errorf("UpdateOrDelete() Run error: %+v\r\n", err)
	}
	fmt.Printf("UpdateOrDelete data id is: %v\r\n", id)
	defer mysqlC.DBClose()
	return
}

// Now
func Test_mysqlDB_Query(t *testing.T) {
	mysqlC, err := New(baseConfig)
	if err != nil {
		t.Errorf("New() Run error: %+v\r\n", err)
	}
	var sqlCommand = "SELECT * FROM pdns.records"
	//var par = []interface{}{"pdns.records"}
	data, err := mysqlC.Query(sqlCommand)
	if err != nil {
		t.Errorf("UpdateOrDelete() Run error: %+v\r\n", err)
	}
	fmt.Printf("UpdateOrDelete data id is: %v\r\n", data)
	defer mysqlC.DBClose()
	return
}

func Test_mysqlDB_QueryRow(t *testing.T) {
	mysqlC, err := New(baseConfig)
	if err != nil {
		t.Errorf("New() Run error: %+v\r\n", err)
	}
	var sqlCommand = "SELECT * FROM pdns.records WHERE id = ?"
	var par = []interface{}{"1"}
	row, err := mysqlC.QueryRow(sqlCommand, par...)
	if err != nil {
		t.Errorf("UpdateOrDelete() Run error: %+v\r\n", err)
	}
	fmt.Printf("UpdateOrDelete data id is: %v\r\n", row)
	defer mysqlC.DBClose()
	return
}

func Test_mysqlDB_Delete(t *testing.T) {
	mysqlC, err := New(baseConfig)
	if err != nil {
		t.Errorf("New() Run error: %+v\r\n", err)
	}
	var sqlCommand = "DELETE FROM records WHERE id = ?"
	var par = []interface{}{"3"}
	id, err := mysqlC.UpdateOrDelete(sqlCommand, par...)
	if err != nil {
		t.Errorf("UpdateOrDelete() Run error: %+v\r\n", err)
	}
	fmt.Printf("UpdateOrDelete data id is: %v\r\n", id)
	defer mysqlC.DBClose()
	return
}
