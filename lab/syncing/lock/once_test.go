package lock

import "sync"

func (s *TrySuite) TestOnce() {
	go doprint()
	go doprint()
}

var a string
var once sync.Once

func setup() {
	a = "hello, world\n"
}

func doprint() {
	once.Do(setup)
	print(a)
}
