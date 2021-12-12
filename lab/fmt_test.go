package lab

import (
	"fmt"
)

type st struct {
	age  int
	name string
}

func (s *TrySuite) TestFmt() {
	b1 := true
	fmt.Printf("%T value: %t %v %+v  %#v\n", b1, b1, b1, b1, b1)

	st1 := &st{
		age:  12,
		name: "boy",
	}
	fmt.Printf("%T value: %v, %+v %#v\n", st1, st1, st1, st1)

	s1 := "hello 曹 go"
	fmt.Printf("%T value: %q %#q %x % x\n", s1, s1, s1, s1, s1)

	i1 := 26361
	fmt.Printf("%T value: %q %#q %x % x\n", i1, i1, i1, i1, i1)

	ui1 := 26361
	fmt.Printf("%T value: %q %#q %x % x\n", ui1, ui1, ui1, ui1, ui1)

	r1 := '曹'
	fmt.Printf("%T value: %q %#q %x % x\n", r1, r1, r1, r1, r1)

	m1 := map[string]int{"cao": 12, "wang": 34}
	fmt.Printf("%T value: %q %#q %x % x\n", m1, m1, m1, m1, m1)

	f1 := 1.234
	fmt.Printf("%T value: %g %x % x\n", f1, f1, f1, f1)
}
