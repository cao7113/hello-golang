package dingtalk
//go:generate mockgen -destination=./mock/dingtalk_sender_mock.go github.com/cao7113/hellogolang/chatops/dingtalk Sender
import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/cao7113/hellogolang/config"
	"github.com/sirupsen/logrus"
)

type Sender interface {
	SendRequest([]byte) error
}

type DingTalkSender struct {
}

func SendRequest(msg []byte) error {
	r := bytes.NewReader(msg)
	url := config.Config.DingdingURL + "?access_token=" + config.Config.DingdingToken
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		logrus.Fatal("DingTalk new request error:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal("DingTalk do request error:", err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Info("DingTalk response body:", string(body))
	return nil
}
