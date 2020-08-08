package model

import (
	"time"

	"github.com/cao7113/hellogolang/config"

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
	config.Conn.Model(&User{}).Count(&cnt)
	return cnt
}

// Create a user
func (u *User) Create() {
	config.Conn.Create(u)
}
