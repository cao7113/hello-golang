package main

import "fmt"

func main() {
	defer fixPanic()
	panic("main panic occur")
}

func fixPanic() {
	if e := recover(); e != nil {
		fmt.Printf("recover panic error: %v\n", e)
	}
	println("trying to fix logic")
}
