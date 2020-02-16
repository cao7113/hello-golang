package pbaccount

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAccount(t *testing.T) {
	tmail := "t@mail.com"
	account := &Account{
		Name:  "tname",
		Email: tmail,
	}
	assert.Equal(t, account.GetEmail(), tmail)
}
