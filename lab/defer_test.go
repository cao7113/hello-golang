package lab

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex
var balance int

func deposit1(amount int) int {
	mu.Lock()
	defer func() {
		fmt.Printf("defer called from deposit1\n")
		mu.Unlock()
	}()
	defer func() {
		fmt.Printf("defer2 called from deposit1 \n")
	}()
	balance += amount
	fmt.Printf("%d is added into balance\n", amount)
	//panic("mock panic from deposit")
	return balance
}

func GetBalance() string {
	mu.Lock()
	defer func() {
		fmt.Printf("defer called from GetBalance \n")
		mu.Unlock()
	}()
	fmt.Printf("before return balance: %d\n", balance)
	return fmt.Sprintf("balance: %d", balance)
}

func TestDefer(t *testing.T) {
	b := deposit1(2)
	fmt.Println(b)
	fmt.Println()
	fmt.Println(GetBalance())
}
