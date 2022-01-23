package lab

import "testing"

// https://github.com/google/pprof/blob/master/doc/README.md#interpreting-the-callgraph
//go:generate go test -bench=. bench_test.go -cpuprofile=pprof-cpu.out -memprofile=pprof-mem.out
//go:generate go tool pprof pprof-cpu.out
func BenchmarkLoopN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopN(40)
	}
}

func BenchmarkFactorialN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialN(40)
	}
}

func FactorialN(n int) int {
	if n > 0 {
		return n * FactorialN(n-1)
	}
	return 1
}

func LoopN(n int) int {
	result := 1
	for i := 0; i <= n; i++ {
		result *= i
	}
	return result
}
