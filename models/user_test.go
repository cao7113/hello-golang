package models

import (
	"testing"
	"time"

	// "github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
)

func TestUserCount(t *testing.T) {
	cnt1 := UsersCount()
	user := &User{
		Name:  "test1",
		Email: time.Now().String(),
	}
	user.Create()
	cnt2 := UsersCount()
	assert.Equal(t, cnt1+1, cnt2)
}
