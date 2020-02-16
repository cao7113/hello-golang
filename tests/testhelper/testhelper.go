package testhelper

import (
	"fmt"

	"github.com/cao7113/hellogolang/database"
)

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate %s;", tableName)
	database.Conn.Exec(sql)
}

func InsertRecord(record interface{}) error {
	err := database.Conn.Create(record).Error
	return err
}
