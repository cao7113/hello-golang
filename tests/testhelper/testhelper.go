package testhelper

import (
	"fmt"
	"github.com/cao7113/hellogolang/db"
)

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate table %s;", tableName)
	db.MyConn.Exec(sql)
}

func InsertRecord(record interface{}) error {
	err := db.MyConn.Create(record).Error
	return err
}

func RecordCount(record interface{}) int64 {
	var count int64
	db.MyConn.Model(record).Count(&count)
	return count
}
