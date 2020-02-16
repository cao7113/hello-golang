package pbuser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUser(t *testing.T) {
	tmail := "t@mail.com"
	user := &User{
		Name:  "tname",
		Email: tmail,
	}
	assert.Equal(t, user.GetEmail(), tmail)
}
