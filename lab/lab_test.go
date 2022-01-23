package lab

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *LabSuite) TestCopy() {
	var s1 []int
	s2 := []int{3, 4}
	copy(s1, s2)
	s.Nil(s1)
	s.Equal(0, len(s1))

	s3 := make([]int, len(s2))
	copy(s3, s2)
	s.Equal(len(s2), len(s3))
	s.Equal(s2, s3)
}

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int) // Create A new channel.
	go generate(ch)      // Start generate() as A subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		if prime > 1000 {
			return
		}
	}
}

func (s *LabSuite) TestSieve() {
	sieve()
}

var a = 1 // file block

func (s *LabSuite) TestBlock() {
	s.EqualValues(1, a)
	var a = 2 // method block
	if true {
		var a = 3 // if block
		s.EqualValues(3, a)
	}
	s.EqualValues(2, a)
}

func TestLabSuite(t *testing.T) {
	suite.Run(t, &LabSuite{})
}

type LabSuite struct {
	suite.Suite
}
