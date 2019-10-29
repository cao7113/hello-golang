package db

import (
	"github.com/cao7113/golang/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbURL = config.Settings.DbURL

var myDb interface{}

func init() {
	myDb, err := gorm.Open("mysql", dbURL)
	defer myDb.Close()
	if err != nil {
		// todo
	}
}
