package errhandling

import "testing"

func Test_hasError(t *testing.T) {
	err := AKindError()
	t.Log(err)
	a, ok := err.(KindError)
	t.Log(a, ok)
}
