package syncing

import (
	"golang.org/x/sync/errgroup"
)

func (s *SyncSuite) TestErrGroup() {
	ch := make(chan int, 5)
	eg := errgroup.Group{}

	eg.Go(func() error { // sender
		for i := 0; i < 10; i++ {
			ch <- i
		}
		println("channel closing")
		close(ch)
		println("sender over")
		return nil
	})

	// listeners
	for i := 0; i < 3; i++ {
		eg.Go(func() error {
			for n := range ch {
				if n%3 == 0 {
					println("listener found num: ", n)
				}
			}
			println("listener end")
			return nil
		})
	}

	err := eg.Wait()
	s.Nil(err)
}
