package dingtalk

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Markdown struct {
	MsgType  string      `json:"msgtype"`
	Markdown MarkdownMsg `json:"markdown"`
	At       At          `json:"at"`
}

type MarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// SendMarkdown: 在 text 内容里要有@手机号, mobiles 才奏效
func SendMarkdown(title string, text string, isAtAll bool, mobiles []string) error {
	msg := Markdown{
		MsgType: "markdown",
		Markdown: MarkdownMsg{
			Title: title,
			Text:  text,
		},
		At: At{
			AtMobiles: mobiles,
			IsAtAll:   isAtAll,
		},
	}

	msgValue, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("DingTalk Markdown json marshal error:", err)
		return err
	}

	return sendRequest(msgValue)
}
