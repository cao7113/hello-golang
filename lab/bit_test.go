package lab

import "fmt"

//https://yourbasic.org/golang/bitmask-flag-set-clear/

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func (s *LabSuite) TestBits() {
	var b Bits
	b = Set(b, F0)
	b = Toggle(b, F2)
	for i, flag := range []Bits{F0, F1, F2} {
		fmt.Println(i, Has(b, flag))
	}
}

func (s *LabSuite) TestOps() {
	b := 0b0010 // 2
	s.Equal(0b0001, b>>1)
	s.Equal(1, b>>1) // divide by 2
	s.Equal(0b0100, b<<1)
	s.Equal(4, b<<1) // multiply by 2
}

func (s *LabSuite) TestPowerOf2() {
	s.True(isPowerOf2(4))
	s.False(isPowerOf2(3))
}

// https://stackoverflow.com/questions/600293/how-to-check-if-a-number-is-a-power-of-2
//  100...000
func isPowerOf2(x uint64) bool {
	return x > 0 && (x&(x-1) == 0)
}

func (s *LabSuite) TestTwoComplement() {
	// https://en.wikipedia.org/wiki/Two%27s_complement
	var n int8 = -1
	fmt.Printf("%0b\n", byte(n))

	n1 := 0b1111_1111 // 1的补码形式， -1， a的补码是-A
	fmt.Printf("%0b\n", byte(n1))
	s.Equal(int8(-1), int8(n1))
}

func (s *LabSuite) TestInt() {
	//const Huge = 1 << 100
	//println(Huge)
	//s.Panics(func() {
	//	const i = uint64(-1)
	//})

	i := ^1
	s.EqualValues(-2, i)
}
