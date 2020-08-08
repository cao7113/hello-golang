package testing

import (
	"fmt"

	"github.com/cao7113/hellogolang/config"
)

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate table %s;", tableName)
	config.Conn.Exec(sql)
}

func InsertRecord(record interface{}) error {
	err := config.Conn.Create(record).Error
	return err
}

func RecordCount(record interface{}) int {
	var count int
	config.Conn.Model(record).Count(&count)
	return count
}
