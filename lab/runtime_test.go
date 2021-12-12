package lab

import (
	"runtime"
	"testing"
)

func TestMaxProcs(t *testing.T) {
	nc := runtime.NumCPU()
	println(nc)
}
