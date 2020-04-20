package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	fmt.Println("teller is started")
	for {
		select {
		case amount := <-deposits:
			fmt.Printf("deposit amount=%d from deposits chan \n", amount)
			balance += amount
		case balances <- balance:
			fmt.Printf("send balance %d to balances chan \n", balance)
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
