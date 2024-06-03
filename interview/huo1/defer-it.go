package main

import "fmt"

//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/

func main() {
	defer func(n int) {
		println("defer: declare before recover-defer ", n)
	}(1)

	//defer recoverv2()
	defer recoverv1()

	defer func(n int) {
		println("defer: declare after recover-defer ", n)
	}(3)

	println("before main panic")
	panic("main panic")
	println("after main panic")
}

func recoverv2() {
	println("in recover-v2")
	recoverv1()
}

func recoverv1() {
	println("in recover-v1")
	// recover 仅在defer中直接-调用才有效！！！ 为什么？
	r := recover()
	if r != nil {
		//log.Infof("error occured: %s", debug.Stack())
		fmt.Printf("try to recover with return info: %v \n", r)
	}
}

/*
// The recover built-in function allows a program to manage behavior of a
// panicking goroutine. Executing a call to recover inside a deferred
// function (but not any function called by it) stops the panicking sequence
// by restoring normal execution and retrieves the error value passed to the
// call of panic. If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the goroutine is not
// panicking, or if the argument supplied to panic was nil, recover returns
// nil. Thus the return value from recover reports whether the goroutine is
// panicking.
func recover() interface{}
*/

// https://go.dev/blog/defer-panic-and-recover#:~:text=Panic%20is%20a%20built%2Din,of%20control%20and%20begins%20panicking.&text=Recover%20is%20a%20built%2Din,and%20have%20no%20other%20effect.
