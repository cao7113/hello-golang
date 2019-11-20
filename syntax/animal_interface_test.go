package syntax

import (
	"fmt"
	"testing"
)

func TestADog(t *testing.T) {
	var animal Animal

	// Interface values with nil underlying values
	fmt.Println("animal value is:", animal)    // animal value is: <nil>
	fmt.Printf("animal type is: %T\n", animal) // animal type is: <nil>

	animal = Dog{"旺财"}
	animal.Bark() // 旺财: wan wan wan!

	fmt.Println("animal value is:", animal)    // animal value is: {旺财}
	fmt.Printf("animal type is: %T\n", animal) // animal type is: Dog
}
