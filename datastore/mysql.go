package datastore

import (
	"log"

	"github.com/cao7113/golang/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MyDB mysql db connection
var MyDB *gorm.DB

// SetupMysql setup conn
func SetupMysql() {
	dsn := config.Settings.DbURL
	MyDB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer MyDB.Close()
}
