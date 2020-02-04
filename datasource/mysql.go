package datasource

import (
	"time"

	"github.com/cao7113/hellogolang/config"
	"github.com/jinzhu/gorm"

	// nolint
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MyDB mysql db connection
var MyDB *gorm.DB

// SetupMysql setup conn
func init() {
	log := config.Logger
	dsn := config.Settings.DbURL
	var err error
	MyDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// defer MyDB.Close()

	MyDB.DB().SetMaxOpenConns(300)
	MyDB.DB().SetMaxIdleConns(100)
	MyDB.DB().SetConnMaxLifetime(5 * time.Minute)

	// Logger.Debugf("db: %+v", MyDB)
}
