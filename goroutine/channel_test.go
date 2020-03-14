package goroutine

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T){
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}