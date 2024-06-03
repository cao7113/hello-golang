package main

import (
	"time"
)

func main() {
	defer println("defer in main") // not run here!!!
	go func() {
		defer println("run defer in go routine")
		panic("go-panic")
		//runtime.Goexit()
	}()

	println("run in main")
	time.Sleep(1 * time.Second)
}
