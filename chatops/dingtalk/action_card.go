package dingtalk

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// ActionCard struct for sending actioncard through DingTalk. Doc: https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1
type ActionCard struct {
	MsgType    string        `json:"msgtype"`
	ActionCard ActionCardMsg `json:"actionCard"`
}

type ActionCardMsg struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	Btns           []Btn  `json:"btns"`
	BtnOrientation string `json:"btnOrientation"`
	HideAvatar     string `json:"hideAvatar"`
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

func SendActionCard(title string, text string, btns []Btn) error {
	msg := ActionCard{
		MsgType: "actionCard",
		ActionCard: ActionCardMsg{
			Title:          title,
			Text:           text,
			Btns:           btns,
			BtnOrientation: "1",
			HideAvatar:     "0",
		},
	}

	msgValue, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("DingTalk ActionCard json marshal error:", err)
		return err
	}

	return SendRequest(msgValue)
}
