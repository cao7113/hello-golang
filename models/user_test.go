package models

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func (s *UserSuite) TestUserCount() {
	cnt1 := UsersCount()
	fk := time.Now().Format(time.RFC3339)
	user := &User{
		Name:  "name-" + fk,
		Email: fk + "@test.com",
	}
	user.Create()

	cnt2 := UsersCount()
	s.Equal(cnt1+1, cnt2)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, &UserSuite{})
}

type UserSuite struct {
	suite.Suite
}
