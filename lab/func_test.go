package lab

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

// https://lessisbetter.site/2019/06/09/golang-first-class-function/#%E7%89%88%E6%9C%AC3

type Checker struct {
	Name       string
	HandleFunc func(age int) string
}

type hook func()

func withHook(name string, h hook) {
	println("before hook")
	if h != nil {
		h()
	}
	println("after hook")
}

func (s *FuncSuite) TestHookd() {
	withHook("no-hook", nil)
	withHook("A-hook", func() {
		println("run A-hook")
	})
	withHook("no-hook-again", nil)
}

func (s *FuncSuite) TestStructFuncArgs() {
	c := &Checker{
		Name: "test-checker",
		HandleFunc: func(age int) string {
			return fmt.Sprintf("age=%d", age)
		},
	}
	s.EqualValues("age=3", c.HandleFunc(3))
}

func TestFuncSuite(t *testing.T) {
	suite.Run(t, &FuncSuite{})
}

type FuncSuite struct {
	suite.Suite
}
