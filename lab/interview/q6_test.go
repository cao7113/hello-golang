package interview

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	panic("panic!")
	//fmt.Printf("hello defer") // never reach
}
