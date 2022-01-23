package painkiller

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *PillSuite) TestConst() {
	s.EqualValues(2, D)
	s.Equal("C", D.String())
	s.Equal("C", E.String())
	s.EqualValues(2, E)
	s.EqualValues(2, F)
}

func TestPillSuite(t *testing.T) {
	suite.Run(t, &PillSuite{})
}

type PillSuite struct {
	suite.Suite
}
