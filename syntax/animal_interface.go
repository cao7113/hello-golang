package syntax

import (
	"fmt"
)

// Animal base interface
type Animal interface {
	Bark() string
}

// Dog an implementation
type Dog struct {
	name string
}

// Bark implement interface methods
func (dog Dog) Bark() string {
	fmt.Println(dog.name + ": wan wan wan!")
	return "wan wan wan"
}
