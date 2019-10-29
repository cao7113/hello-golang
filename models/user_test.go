package models

import (
	"log"
	"testing"

	"github.com/cao7113/golang/config"
	"github.com/cao7113/golang/datastore/mysql"
	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
)

func TestUserCount(t *testing.T) {
	dsn := config.Settings.DbURL
	c, err := gorm.Open("mysql", dsn)
	defer c.Close()
	if err != nil {
		log.Fatalln(err)
	}

	cnt := 0
	c.Model(&User{}).Count(&cnt)
	assert.Equal(t, 0, cnt)
}

func TestUserCount1(t *testing.T) {

	mysql.Setup()
	cnt1 := UsersCount()

	assert.Equal(t, 0, cnt1)
	// user := &User{
	// 	Name: "test1",
	// }
	// user.Create()
	cnt2 := UsersCount()
	assert.Equal(t, cnt1+1, cnt2)

}
