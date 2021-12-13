package db

import (
	"github.com/cao7113/hellogolang/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var MyConn *gorm.DB

func init() {
	dsn := config.Config.DbURL
	if strings.TrimSpace(dsn) == "" {
		panic("no mysql dsn configured")
	}

	var err error
	MyConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open mysql error")
	}
	// defer Conn.Close()

	sqlDB, err := MyConn.DB()
	if err != nil {
		panic("get sqlDB error")
	}
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
