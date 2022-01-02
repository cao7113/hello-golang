//go:build ignore

package buildtag

import log "github.com/sirupsen/logrus"

func (s *TagSuite) TestIgnore() {
	log.Info("hi ignore tag")
	s.Equal(IgnoredWord, IgnoreString())
}
