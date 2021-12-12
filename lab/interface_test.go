package lab

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func (s *InterfaceSuite) TestADog() {
	var animal Animal

	s.Nil(animal)                                     // animal value is: <nil>
	s.EqualValues("<nil>", fmt.Sprintf("%T", animal)) // animal type is: <nil>

	animal = Dog{"旺财"}
	s.EqualValues("wang wang wang!!!", animal.Bark())

	fmt.Println("animal value is:", animal)    // animal value is: {旺财}
	fmt.Printf("animal type is: %T\n", animal) // animal type is: Dog
}

func TestInterfaceSuite(t *testing.T) {
	suite.Run(t, &InterfaceSuite{})
}

type InterfaceSuite struct {
	suite.Suite
}

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
	log.Println(dog.name, ": wan wan wan!")
	return "wang wang wang!!!"
}
