package muser

import (
	"github.com/cao7113/hellogolang/testing/trymock/mail"
	"github.com/sirupsen/logrus"
)

type Muser struct {
	Email string
	Name  string
}

func (m Muser) WelcomeBy(sender mail.Sender) string {
	logrus.Infof("sent welcome mail to :%s ", m.Email)
	return sender.Send(m.Email)
}
