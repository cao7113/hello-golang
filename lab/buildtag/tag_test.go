package buildtag

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *TagSuite) TestIt() {
	fmt.Println("hi build-tag")
}

func TestTagSuite(t *testing.T) {
	suite.Run(t, &TagSuite{})
}

type TagSuite struct {
	suite.Suite
}
