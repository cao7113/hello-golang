package interview

import "testing"

type Saying interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestAssign(t *testing.T) {
	//var stu Saying = Student{}
}
