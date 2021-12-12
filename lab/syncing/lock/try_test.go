package lock

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var i int

func (s *TrySuite) TestGoRoutines() {
	go add(&i)
	go add(&i)

	time.Sleep(time.Second * 3)

	println(i)
}

func add(i *int) {
	for j := 0; j < 10000; j++ {
		*i = *i + 1
	}
}

func (s *TrySuite) TestOperateSameVar() {
	go func() {
		s.Equal(1, s.num)
		s.num = 2
		//time.Sleep(1 * time.Second)
	}()
	time.Sleep(10 * time.Millisecond)
	s.Equal(2, s.num)
}

func TestTrySuite(t *testing.T) {
	suite.Run(t, &TrySuite{})
}

type TrySuite struct {
	suite.Suite
	num int
}
