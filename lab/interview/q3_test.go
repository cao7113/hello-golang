package interview

import (
	"fmt"
	"testing"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestShowA(tt *testing.T) {
	t := Teacher{}
	t.ShowA()
}
