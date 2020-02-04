package models

import (
	"time"

	ds "github.com/cao7113/hellogolang/datasource"

	"github.com/jinzhu/gorm"
)

// User user
type User struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
}

// UsersCount get users count
func UsersCount() int {
	var cnt int
	ds.MyDB.Model(&User{}).Count(&cnt)
	return cnt
}

// Create a user
func (u *User) Create() {
	ds.MyDB.Create(u)
}
