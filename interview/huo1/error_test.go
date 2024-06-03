package main

import "fmt"

type SampleError int

func (s *SampleError) Error() string {
	return ""
}

func (s *ASuite) TestNilErr() {
	e := foo()
	fmt.Println(e)
	s.False(e == nil)
}

func foo() error {
	var err *SampleError
	return err
}
