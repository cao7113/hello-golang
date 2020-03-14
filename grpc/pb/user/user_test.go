package pbuser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUser(t *testing.T) {
	tmail := "t@mail.com"
	user := &User{
		Email: tmail,
	}
	assert.Equal(t, user.GetEmail(), tmail)
	assert.Equal(t, user.GetName(), "")
	assert.Equal(t, user.GetIsBoy(), false)
	if user.GetIsGirl() != nil {
		t.Error("is_girl should nil")
	}
	if user.GetJob() != nil {
		t.Error("job should nil")
	}
	if user.GetUpdatedAt() != nil {
		t.Error("updated_at should nil")
	}
}
