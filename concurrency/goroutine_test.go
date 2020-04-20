package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutine1(t *testing.T) {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	time.Sleep(time.Second)
}

//  说出可能执行结果 https://docs.hacknode.org/gopl-zh/ch9/ch9-04.html
