package muser

import (
	"testing"

	mock_mail "github.com/cao7113/hellogolang/testing/trymock/mail/mock"

	"github.com/magiconair/properties/assert"

	"github.com/golang/mock/gomock"
)

func TestMuser_WelcomeBy(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockSender := mock_mail.NewMockSender(ctl)
	gomock.InOrder(
		mockSender.EXPECT().Send("a@b.c").Return("sent to a@b.c"),
	)
	mu := &Muser{Email: "a@b.c"}
	resp := mu.WelcomeBy(mockSender)
	assert.Equal(t, "sent to a@b.c", resp)
}
