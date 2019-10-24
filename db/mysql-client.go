package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
const dbURL = "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"

var myDb interface{}

func init() {
	myDb, err := gorm.Open("mysql", dbURL)
	defer myDb.Close()
	if err != nil {
		// todo
	}
}
