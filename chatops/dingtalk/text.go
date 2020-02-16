package dingtalk

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Text struct {
	MsgType string  `json:"msgtype"`
	Text    TextMsg `json:"text"`
	At      At      `json:"at"`
}

type TextMsg struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

func SendText(content string, isAtAll bool, mobiles []string) error {
	msg := Text{
		MsgType: "text",
		Text: TextMsg{
			Content: content,
		},
		At: At{
			AtMobiles: mobiles,
			IsAtAll:   isAtAll,
		},
	}

	msgValue, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("DingTalk Text json marshal error:", err)
		return err
	}

	return sendRequest(msgValue)
}
