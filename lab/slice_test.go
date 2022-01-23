package lab

import (
	"encoding/binary"
	"fmt"
)

//https://go.dev/blog/slices-intro // 讲解非常深入
// Array is value, slice is reference

func (s *LabSuite) TestSliceRef() {
	sl := []int{3, 2, 1}
	sn := sl[1:] // just ref raw underlying slice
	sn[0] = 4    // change raw array element
	s.Equal(4, sl[1])
	s.Equal(4, sn[0])
	s.Equal(2, len(sn))
	s.Equal(2, cap(sn))

	sn = append(sn, 8) // trigger re-slice to new underlying array
	sn[0] = 6
	s.Equal(4, sl[1]) // old array slice
	s.Equal(6, sn[0]) // new gen-array slice

	s.Equal(3, len(sn))
	s.Equal(4, cap(sn)) // slice re-slicing to new size, 2**B
}

func (s *LabSuite) TestBytes() {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 2)
}

func (s *LabSuite) TestMake() {
	l := make([]int, 3, 5)
	fmt.Printf("%T %v cap: %d", l, l, cap(l))
	l[2] = 2
	fmt.Printf("%T %v", l, l)

	nl := l[:]
	fmt.Printf("%T %v", nl, nl)

	nl = append(nl, 4)
	fmt.Printf("%T %v", nl, nl)

	//panic("A")
}

func (s *LabSuite) TestArrayAndSlice() {
	intSet := [6]int{1, 2, 3, 5}
	days := [...]string{"Sat", "Sun"} // ... = max-index + 1
	s.Equal(6, len(intSet))
	s.Equal(2, len(days))

	// list of prime numbers
	primes := []int{2, 3, 5, 7, 9, 2147483647}
	// vowels[ch] is true if ch is A vowel
	vowels := [128]bool{'A': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
	// the array [10]float32{-1, 0, 0, 0, -0.1, -0.2, 0, 0, 0, -1}
	filter := [10]float32{-1, 4: -0.1, -0.2, 9: -1}

	s.Equal(3, primes[1])
	s.True(vowels['e'])
	s.EqualValues(-0.2, filter[5])
}

func (s *LabSuite) TestVarArgs() {
	sl := []int{1, 2, 3}
	sum := varArgs(sl...) // 共享底层数组
	s.Equal(6, sum)
	s.Equal([]int{2, 3, 4}, sl)
}

func varArgs(items ...int) int {
	var sum int
	for i, it := range items {
		items[i] = it + 1 // try to change slice
		sum += it
	}
	return sum
}
