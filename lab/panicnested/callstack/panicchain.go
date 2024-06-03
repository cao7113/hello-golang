package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("main recover error: %v\n", r)
		}
	}()
	println("main run")
	foo()
	time.Sleep(1 * time.Second)
}

func foo() {
	defer func() {
		println("foo defer run")
	}()
	panicFoo()
}

func panicFoo() {
	panic("foo panic")
}
