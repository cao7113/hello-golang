package interview

import (
	"fmt"
	"testing"
)

func TestQ1s1(t *testing.T) {
	c := make(chan string, 10)
	close(c)
	c <- "test string"
}

func TestQ1s2(t *testing.T) {
	c := make(chan string, 10)
	close(c)
	for s := range c {
		fmt.Printf("received: %s \n", s)
	}
}

func TestQ1s3(t *testing.T) {
	c := make(chan string, 10)
	close(c)
	for {
		select {
		case s := <-c:
			fmt.Printf("received: %s \n", s)
		default:
			fmt.Println("no data")
		}
	}
}
