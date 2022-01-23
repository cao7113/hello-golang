package lab

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *SelectSuite) TestSelect() {
	//s.Run("blank select", func() {
	//	select {} // hang here
	//	// never here
	//})

	s.Run("basic", func() {
		var c1, c2, c3 chan int
		var i1, i2 int
		select {
		case i1 = <-c1:
			fmt.Println("received ", i1, " from c1")
		case c2 <- i2:
			fmt.Println("sent ", i2, " to c2")
		case i3, ok := <-c3:
			if ok {
				fmt.Println("received ", i3, " from c3")
			} else {
				fmt.Println("c3 is closed")
			}
		default:
			fmt.Println("no communication")
		}
	})
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, &SelectSuite{})
}

type SelectSuite struct {
	suite.Suite
}

func f() int {
	return 0
}

func (s *LabSuite) TestSelect() {
	var a []int
	var c, c1, c2, c3, c4 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		print("received ", i1, " from c1\n")
	case c2 <- i2:
		print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			print("received ", i3, " from c3\n")
		} else {
			print("c3 is closed\n")
		}
	case a[f()] = <-c4:
		// same as:
		// case t := <-c4
		//	A[f()] = t
	default:
		print("no communication\n")
	}

	for { // send random sequence of bits to c
		select {
		case c <- 0: // note: no statement, no fallthrough, no folding of cases
		case c <- 1:
		}
	}

	//select {} // block forever // unreachable: unreachable code
}
