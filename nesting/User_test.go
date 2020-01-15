package nesting

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestingJson(t *testing.T) {
	buser := &UserBlog{
		User: User{
			Name:  "a",
			Email: "b@c.com",
		},
		Title: "test",
	}
	s, err := json.Marshal(buser)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("json format: %+v\n", string(s))
	jstr := `{"name":"a","email":"b@c.com","title":"test"}`
	assert.Equal(t, string(s), jstr)

	buser1 := &UserBlog{}
	json.Unmarshal([]byte(jstr), buser1)
	assert.Equal(t, buser1.Name, "a")
	assert.Equal(t, buser1.Title, "test")
	// &{User:{Name:a Email:b@c.com} Title:test}
	fmt.Printf("object format: %+v\n", buser1)
}

// json:"email,omitempty"`
func TestOmitJson(t *testing.T) {
	buser := &UserBlog{
		User: User{
			Name: "a",
		},
		Title: "test",
	}
	s, err := json.Marshal(buser)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("json format: %+v\n", string(s))
	jstr := `{"name":"a","title":"test"}`
	assert.Equal(t, string(s), jstr)

	buser = &UserBlog{
		Title: "test",
	}
	s, err = json.Marshal(buser)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("json format: %+v\n", string(s))
	jstr = `{"name":"","title":"test"}`
	assert.Equal(t, string(s), jstr)
}
