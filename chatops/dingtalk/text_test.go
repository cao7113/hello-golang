package dingtalk

import "testing"

// https://ding-doc.dingtalk.com/doc#/serverapi3/iydd5h
func TestDingtalkNotifier_Send(t *testing.T) {
	ats := []string{}
	SendWithTitle("Up: 天天开心", "今天你🙂了吗？", ats)
}
