package ctx

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

// https://blog.golang.org/context

func (s *ContextSuite) TestCancel() {
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)
	s.NotNil(ctx.Done())
	s.Nil(ctx.Err())

	// wait and cancel
	go func() {
		log.Println("[watcher] cancel called", <-time.After(3*time.Second))
		cancel()
	}()

	// read numbers
	out := make(chan int)
	go func() {
		for i := range out {
			log.Println("[reader] got number: ", i)
		}
	}()

	for {
		if ctx.Err() != nil {
			log.Println("hit ctx err: ", ctx.Err())
			break
		}

		select {
		case <-ctx.Done():
			log.Errorf("done error: %s", ctx.Err())
		case out <- genNumber():
		}
	}
}

func genNumber() int {
	rand.Seed(time.Now().Unix())
	time.Sleep(1 * time.Second)
	num := rand.Int()
	return num
}

func TestContextSuite(t *testing.T) {
	suite.Run(t, &ContextSuite{})
}

type ContextSuite struct {
	suite.Suite
}
