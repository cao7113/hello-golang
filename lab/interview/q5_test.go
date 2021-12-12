package interview

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func TestFor(t *testing.T) {
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	m := make(map[string]*student)
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Printf("%s: %+v\n", k, v)
	}
}
