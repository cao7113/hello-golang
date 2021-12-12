package lab

import "fmt"

//https://go.dev/blog/slices-intro // 讲解非常深入
// Array is value, slice is reference

func (s *TrySuite) TestMake() {
	l := make([]int, 3, 5)
	fmt.Printf("%T %v cap: %d", l, l, cap(l))
	l[2] = 2
	fmt.Printf("%T %v", l, l)

	nl := l[:]
	fmt.Printf("%T %v", nl, nl)

	nl = append(nl, 4)
	fmt.Printf("%T %v", nl, nl)

	panic("a")
}

func (s *TrySuite) TestArrayAndSlice() {
	intSet := [6]int{1, 2, 3, 5}
	days := [...]string{"Sat", "Sun"} // ... = max-index + 1
	s.Equal(6, len(intSet))
	s.Equal(2, len(days))

	// list of prime numbers
	primes := []int{2, 3, 5, 7, 9, 2147483647}
	// vowels[ch] is true if ch is a vowel
	vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
	// the array [10]float32{-1, 0, 0, 0, -0.1, -0.2, 0, 0, 0, -1}
	filter := [10]float32{-1, 4: -0.1, -0.2, 9: -1}

	s.Equal(3, primes[1])
	s.True(vowels['e'])
	s.EqualValues(-0.2, filter[5])
}

type Hook func()

func mWithHooks(name string, hooks ...Hook) {
	println(name)
	for _, h := range hooks {
		h()
	}
}

func (s *TrySuite) TestSliceHooks() {
	mWithHooks("hi")
	mWithHooks("f", func() {
		println("f1")
	}, func() {
		println("f2")
	})

	hooks := []Hook{func() {
		println("a1")
	}, func() {
		println("a2")
	}}
	mWithHooks("a", hooks...)
}
