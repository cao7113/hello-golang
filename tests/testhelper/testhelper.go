package testhelper

import (
	"fmt"

	"github.com/cao7113/hellogolang/database"
)

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate table %s;", tableName)
	database.Conn.Exec(sql)
}

func InsertRecord(record interface{}) error {
	err := database.Conn.Create(record).Error
	return err
}

func RecordCount(record interface{}) int {
	var count int
	database.Conn.Model(record).Count(&count)
	return count
}
