package syncing

import (
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

func (s *SyncSuite) TestMutex() {
	var mu sync.Mutex
	i := 3
	mu.Lock()
	i++
	mu.Unlock()
	s.Equal(4, i)
}

func TestSyncSuite(t *testing.T) {
	suite.Run(t, &SyncSuite{})
}

type SyncSuite struct {
	suite.Suite
}
