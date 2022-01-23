package lab

import (
	"math/rand"
	"runtime"
	"sync"
)

func (s *LabSuite) TestGPM() {
	println("num of cpu: ", runtime.NumCPU())
	println("num of P: ", runtime.GOMAXPROCS(-1))
	println("num of g: ", runtime.NumGoroutine())

	wg := sync.WaitGroup{}
	cnt := 10
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			defer wg.Done()
			println(i, "num of g: ", runtime.NumGoroutine())
			//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			seed := rand.Intn(50)
			println(i, "Fibonacci of ", seed, " is ", Fibonacci(seed))
		}(i)
	}
	wg.Wait()
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
