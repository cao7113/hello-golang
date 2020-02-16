package pbblog

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBlog(t *testing.T) {
	title := "test title"
	blog := &Blog{
		Title: title,
		Body:  "test body",
	}
	assert.Equal(t, blog.GetTitle(), title)
}
