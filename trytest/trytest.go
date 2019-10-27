package trytest

import "fmt"

// AddOne add one
func AddOne(num int) (result int) {
	result = num + 1
	return
}

// TryLint try lint
func TryLint() {
	fmt.Println("try GoLint")
	// mock vet error
	// fmt.Printf("%s, %v", "hello")

}
