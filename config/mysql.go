package config

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"

	// nolint
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Conn mysql db connection
var Conn *gorm.DB

// SetupMysql setup conn
func init() {
	dsn := Config.DbURL
	var err error
	Conn, err = gorm.Open("mysql", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}
	// defer Conn.Close()

	Conn.DB().SetMaxOpenConns(300)
	Conn.DB().SetMaxIdleConns(100)
	Conn.DB().SetConnMaxLifetime(5 * time.Minute)
}
