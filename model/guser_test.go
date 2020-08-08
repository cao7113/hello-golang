package model

import (
	"testing"

	"github.com/cao7113/hellogolang/config"

	"github.com/cao7113/hellogolang/testing"

	"github.com/magiconair/properties/assert"
)

// http://gorm.io/zh_CN/docs/transactions.html
func TestTrans(t *testing.T) {
	testing.TruncateTable(Guser{}.TableName())
	err := doInTranc("a@b.c")
	assert.Equal(t, err.Error(), "Error 1062: Duplicate entry 'a@b.c' for key 'idx_email'")
	// count should 0
	n := testing.RecordCount(&Guser{})
	assert.Equal(t, 0, n)

	// case2
	testing.TruncateTable(Guser{}.TableName())
	err = doInTranc("a1@b.c")
	assert.Equal(t, err, nil)
	assert.Equal(t, 2, testing.RecordCount(&Guser{}))
}

func doInTranc(email string) error {
	db := config.Conn

	// 请注意，事务一旦开始，你就应该使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Guser{Email: "a@b.c"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Guser{Email: email}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
