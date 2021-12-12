package timer

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func (s *TimerSuite) TestTimer() {
	s.Run("AfterFunc", func() {
		// do something after some time
		tmr := time.AfterFunc(10*time.Millisecond, func() {
			log.Println("AfterFunc timer end")
		})
		time.Sleep(20 * time.Millisecond)
		s.Nil(tmr.C)        // has no channel allocated
		s.False(tmr.Stop()) // already expired

		// quick to get a timer.C channel
		tm := <-time.After(1 * time.Second)
		log.Println("After timer end at ", tm)
	})

	s.Run("NewTimer", func() {
		tmr := time.NewTimer(1 * time.Second)
		log.Println("NewTimer started")
		// provide a timer.C channel to check when timer end
		if tm := <-tmr.C; !tm.IsZero() {
			log.Println("NewTimer end")
		}
		s.NotNil(tmr.C)     // do-not close channel after timeout
		s.False(tmr.Stop()) // already expired
	})

	s.Run("stop early timer", func() {
		tmr := time.NewTimer(1 * time.Second)
		log.Println("stop-early NewTimer started")
		go func() {
			time.AfterFunc(10*time.Millisecond, func() {
				ok := tmr.Stop() // prevents the Timer from firing
				s.True(ok)       // should success stop timer
				s.NotNil(tmr.C)  // channel not closed
				log.Println("stopped timer")
				if !ok {
					tm := <-tmr.C // drain it
					log.Println("drain timer.C after stop, got time: ", tm)
				}
			})
		}()
		time.Sleep(20 * time.Millisecond)
		log.Println("waiting timer signal")
		select {
		case tm := <-tmr.C: // block to wait data, never come after stopped
			log.Println("got time: ", tm)
		default:
			log.Println("default case")
		}
	})
}

func TestTimerSuite(t *testing.T) {
	suite.Run(t, &TimerSuite{})
}

type TimerSuite struct {
	suite.Suite
}
