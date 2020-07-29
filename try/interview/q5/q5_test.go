package q5

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type student struct {
	Name string
	Age  int
}

func TestA(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	logrus.Infof("m=%+v", m)
}
