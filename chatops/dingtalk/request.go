package dingtalk

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/peatio/butterfly/config"
	log "github.com/sirupsen/logrus"
)

func sendRequest(msg []byte) error {
	r := bytes.NewReader(msg)
	url := config.Config.DingdingURL + "?access_token=" + config.Config.DingdingToken
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		log.Fatal("DingTalk new request error:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("DingTalk do request error:", err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("DingTalk response body:", string(body))
	return nil
}
