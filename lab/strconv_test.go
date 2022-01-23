package lab

import (
	"fmt"
	"strconv"
)

func (s *LabSuite) TestStrConv() {
	a := "abc"
	s.Equal("\"abc\"", strconv.Quote(a))
	println(a)
	println(strconv.Quote(a))

	i := 32
	fmt.Printf("%s", strconv.FormatUint(uint64(i), 2))
}

func (s *LabSuite) TestLower() {
	c := 'x' - 'X'
	fmt.Printf("%+v=>%08b\n x=%x X=%x\n", c, c, 'x', 'X')
}

func (s *LabSuite) TestShiftBits() {
	i := 32 << (^uint(0) >> 63)
	fmt.Printf("%0b\n", i)

	fmt.Printf("%0b\n", uint(0))
	fmt.Printf("%b\n", ^uint(0))
	fmt.Printf("%b\n", ^uint(0)>>63)
	j := ^uint(0) >> 63
	fmt.Printf("%0b\n", 32<<j)

	k, _ := strconv.ParseInt("12", 10, 64)
	s.Equal(12, k)
}
