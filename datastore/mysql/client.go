package mysql

import (
	"log"

	"github.com/cao7113/golang/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Client mysql db connection
var Client *gorm.DB

// Setup setup conn
func Setup() {
	dsn := config.Settings.DbURL
	Client, err := gorm.Open("mysql", dsn)
	defer Client.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
