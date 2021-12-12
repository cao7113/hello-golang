package lab

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

func (s *ReflectSuite) TestReflect() {
	var i interface{}
	var vi *int
	i = vi

	t := reflect.TypeOf(i)
	fmt.Printf("t: %T %+v kind: %+v\n", t, t, t)

	v := reflect.ValueOf(i)
	fmt.Printf("v: %T %+v kind: %+v\n", v, v, v.Kind())

	s.NotEqual(nil, i)
	s.Equal(vi, i)
	s.Nil(i)
	s.Nil(vi)
	s.True(vi == nil)
	//s.False(i == nil) // type not equal
	s.True(i == vi)
}

func (s *ReflectSuite) TestAddressable() {
	a := struct{ a string }{"123"}
	t := reflect.TypeOf(a)
	fmt.Printf("%v", t)

	v := reflect.ValueOf(a)
	s.False(v.CanAddr())
}

func TestReflectSuite(t *testing.T) {
	suite.Run(t, &ReflectSuite{})
}

type ReflectSuite struct {
	suite.Suite
}
