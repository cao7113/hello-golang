package models

import (
	"github.com/cao7113/hellogolang/db"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
}

func UsersCount() int64 {
	var cnt int64
	db.MyConn.Model(User{}).Count(&cnt)
	return cnt
}

func (u *User) Create() *gorm.DB {
	return db.MyConn.Create(u)
}
