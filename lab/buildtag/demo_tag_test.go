//go:build demo

// hi:test

package buildtag

import "fmt"

func (s *TagSuite) TestDemo() {
	fmt.Println("demo tag")
	s.True(IsDemo())
}
