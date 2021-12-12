package main

import (
	"github.com/cao7113/hellogolang/lab/pkg/ap"
	"github.com/cao7113/hellogolang/lab/pkg/bp"
)

// https://raw.githubusercontent.com/yangwenmai/maiyang.me/master/blog/init.png

func init() {
	println("main init")
}

func main() {
	println("main run")
	println(ap.A, ap.A1, bp.B)
}
