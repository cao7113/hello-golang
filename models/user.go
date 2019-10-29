package models

import (
	"time"

	"github.com/cao7113/golang/datastore"
	
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
}

// UsersCount get users count
func UsersCount() int32 {
	var cnt int32
	datastore.MyDB.Model(&User{}).Count(&cnt)
	return cnt
}

// Create a user
func (u *User) Create() {
	datastore.MyDB.Create(u)
}
