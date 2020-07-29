package q2

import "testing"

type People interface {
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

//func main() {
//	var peo People = Student{}
//}

func TestAssign(t *testing.T) {
	var peo People = Student{}
}
