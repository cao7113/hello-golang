package dingtalk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/peatio/butterfly/config"
	"github.com/sirupsen/logrus"
)

type Notifier interface {
	Send(title, msg string, mobiles []string)
}

type DingtalkNotifier struct {
}

type MockNotifier struct {
}

func (n *DingtalkNotifier) Send(title, msg string, mobiles []string) {
	content := fmt.Sprintf("%s\n%s", title, msg)
	err := SendText(content, false, mobiles)
	if err != nil {
		logrus.Error(err)
	}
}

func (n *MockNotifier) Send(title, msg string, mobiles []string) {
	template := `{"title":"%s","msg":%s,"mobiles":%s}`
	serializedMobiles, _ := json.Marshal(mobiles)
	serializedMsg, _ := json.Marshal(msg)
	final := fmt.Sprintf(template, title, serializedMsg, string(serializedMobiles))
	config.L1Cache.Set("dingtalk_msg_unit_test", final, 10*time.Minute)
}

func GetNotifier() Notifier {
	if config.IsTest() || os.Getenv("MOCK_DINGTALK") != "" {
		return &MockNotifier{}
	} else {
		return &DingtalkNotifier{}
	}
}
