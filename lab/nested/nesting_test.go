package nested

import (
	"encoding/json"
	"github.com/cao7113/hellogolang/lab/nested/apkg"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *NestingSuite) TestNestInterface() {
	a := &Chinese{
		Greeter: ZhGreeter(""),
		City:    "Beijing",
	}
	s.EqualValues("", a.Greeter)
	s.EqualValues("你好", a.Hi()) // Method up from Greeter
	s.EqualValues("你好", a.Greeter.Hi())
}

func (s *NestingSuite) TestNestMethods() {
	f := &FruitStore{
		&apkg.Apple{Color: "green"},
		"AppleStore",
	}
	f.SetColor("red")             // call nesting-obj method
	s.EqualValues("red", f.Color) // read nesting-obj attribute
	s.EqualValues("AppleStore", f.Name)
}

func (s *NestingSuite) TestNestJson() {
	blog := &Blog{
		User: User{
			Name:  "a",
			Email: "b@c.com",
		},
		Title: "test",
	}
	bs, err := json.Marshal(blog)
	s.Nil(err)
	js := `{"name":"a","email":"b@c.com","title":"test"}`
	s.EqualValues(js, string(bs))

	s.Equal(5, blog.Score())
	s.Equal(5, blog.User.Score())

	blog1 := &Blog{}
	err = json.Unmarshal([]byte(js), blog1)
	s.Nil(err)
	s.Equal("a", blog1.Name)
	s.Equal("test", blog1.Title)
}

func (s *NestingSuite) TestOmitJson() {
	blog := &Blog{
		User: User{
			Name: "a",
		},
		Title: "test",
	}
	bs, err := json.Marshal(blog)
	s.Nil(err)
	js := `{"name":"a","title":"test"}` // no email part
	s.Equal(js, string(bs))

	blog = &Blog{
		Title: "test",
	}
	bs, err = json.Marshal(blog)
	s.Nil(err)
	js = `{"name":"","title":"test"}` // no email part
	s.Equal(js, string(bs))
}

func TestNestingSuite(t *testing.T) {
	suite.Run(t, &NestingSuite{})
}

type NestingSuite struct {
	suite.Suite
}
