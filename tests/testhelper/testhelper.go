package testhelper

import (
	"fmt"
	"github.com/cao7113/hellogolang/db"
)

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate table %s;", tableName)
	db.Conn.Exec(sql)
}

func InsertRecord(record interface{}) error {
	err := db.Conn.Create(record).Error
	return err
}

func RecordCount(record interface{}) int64 {
	var count int64
	db.Conn.Model(record).Count(&count)
	return count
}
