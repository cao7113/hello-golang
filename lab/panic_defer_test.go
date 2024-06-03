package lab

import "fmt"

func (s *LabSuite) TestX() {
	//*(*int)(nil) = 0
	s.Equal(1, deferOk())

	s.Equal(1, deferArgEva())

	deferOrder()

	s.Equal(2, c())
}

func deferOk() (n int) {
	defer func() {
		println("defer1 run")
		n++
	}()

	return n

	// not reachable
	defer func() {
		println("defer2 run")
		n++
	}()
	return 0
}

func deferArgEva() int {
	i := 0
	defer fmt.Println(i) // print i is 0
	i++
	return i
}

func deferOrder() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func c() (i int) {
	defer func() { i++ }() // actual return 2
	return 1
}
