package syncing

import (
	"sync"
	"time"
)

func (s *SyncSuite) TestOnce() {
	go tryDo()
	go tryDo()
	go tryDo()
	time.Sleep(1 * time.Second)
	s.Equal(1, idx)
}

var idx int
var once sync.Once

func setupInit() {
	println("setupInit: should do only once")
	idx++
}

func tryDo() {
	once.Do(setupInit)
}
