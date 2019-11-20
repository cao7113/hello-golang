package errorhandling

import "testing"

func Test_hasError(t *testing.T) {
	var err error
	err = AKindError()
	t.Log(err)
	a, ok := err.(KindError)
	t.Log(a, ok)
}
