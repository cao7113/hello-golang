package datastore

import (
	"log"
	"fmt"
	"time"

	"github.com/cao7113/golang/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MyDB mysql db connection
var MyDB *gorm.DB

// SetupMysql setup conn
func init() {
	dsn := config.Settings.DbURL
	fmt.Printf("==%s", dsn)
	var err error
	MyDB, err = gorm.Open("mysql", dsn)
	if err != nil {

	fmt.Printf("==00011%s", dsn)
		log.Fatalln(err)
	}
	fmt.Printf("\n==111%s", dsn)
	// defer MyDB.Close()

	MyDB.DB().SetMaxOpenConns(300)
	MyDB.DB().SetMaxIdleConns(100)
	MyDB.DB().SetConnMaxLifetime(5 * time.Minute)

	fmt.Printf("\n==2%+v", MyDB)
}
