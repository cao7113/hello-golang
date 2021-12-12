package channel

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
* 向一个nil channel发送消息，会一直阻塞
* 向一个已经关闭的channel发送消息，会引发运行时恐慌（panic）
* channel关闭后不可以继续向channel发送消息，但可以继续从channel接收消息
* 当channel关闭并且缓冲区为空时，继续从从channel接收消息会得到一个对应类型的零值
 */

func (s *ChannelSuite) TestClose() {
	c := make(chan int)
	close(c)
	for i := range c {
		println(i)
	}
	j := <-c
	s.Equal(0, j)

	s.Panics(func() {
		c <- 3 // has closed
	})

	//cr := make(<-chan int)
	//close(cr)
}

func (s *ChannelSuite) TestChan() {
	s.Run("declare and select", func() {
		var ch chan int
		s.Nil(ch)

		var i int
		select {
		case i = <-ch: // receive operation might block a goroutine because of the 'nil' channel
			log.Println("read i=", i)
		default:
			i = 999
		}
		s.Equal(999, i)
	})

	// Note that it is only necessary to close a channel if the receiver is looking for a close.
	// Closing the channel is a control signal on the channel indicating that no more data follows.
	s.Run("declare and close", func() {
		ch := make(chan int)
		//close(ch) // panic: close of nil channel [recovered]
		close(ch)
		s.NotNil(ch) // close not nil
		//close(ch) // panic: close of closed channel
		j := <-ch // read from closed channel, non-block and get zero value
		s.Equal(0, j)

		i, ok := <-ch // read from closed channel and check whether closed
		s.False(ok)
		s.Equal(0, i)

		for k := range ch { // 关闭channel 不会进入
			log.Fatalf("never run here k = %d ", k)
		}

		log.Println("end")
	})

	//s.Run("closed channel and for-select", func() {
	//	ch := make(chan int)
	//	close(ch)
	//	ch2 := make(chan int)
	//	close(ch2)
	//	i := 0
	//	for { // endless loop
	//		select {
	//		case i1 := <-ch:
	//			log.Println(i, "read from closed channel", i1) // rand run this
	//		case i2 := <-ch2:
	//			log.Println(i, "read from closed channel2", i2) // rand run this
	//		default:
	//			log.Println("default in closed channel select", i) // never run here
	//		}
	//		i++
	//	}
	//})

	s.Run("make and read", func() {
		ch := make(chan int)
		defer close(ch)

		var i int
		select {
		case i = <-ch:
			log.Println("read i=", i)
		default:
			i = 999
		}
		s.Equal(999, i)
	})

	s.Run("write and read", func() {
		ch := make(chan int)
		go func(chan int) {
			defer close(ch)
			for i := 0; i < 3; i++ {
				ch <- i
			}
		}(ch)

		// for range是阻塞式读取channel，只有channel close之后才会结束
		// 如果channel 没有关闭，那么会一直等待下去，出现 deadlock 的错误
		for i := range ch { // will block here if not close ch
			log.Println("i =", i)
		}
		log.Println("end")
	})
}

func (s *ChannelSuite) TestChanBuf() {
	s.Run("no buf", func() {
		ch := make(chan int)
		go func() {
			<-time.After(3 * time.Second)
			ch <- 123
		}()
		log.Println("wait num")
		i := <-ch // block 3s until other goroutine write num to chan
		s.Equal(123, i)
		log.Println("go i=", i)
	})

	s.Run("has buf", func() {
		ch := make(chan int, 1)
		ch <- 123
		//ch <- 124 // writer will block to reader take out num from chan
		i, ok := <-ch
		s.True(ok)
		s.Equal(123, i)
		//<-ch // reader will block to wait next num
	})
}

func (s *ChannelSuite) TestDemo() {
	s.Run("timeout control", func() {
		i := 0
		select {
		case <-doSomething(100 * time.Millisecond):
			i = 1
		case <-time.After(20 * time.Millisecond):
			i = 2
		}
		s.Equal(2, i)
	})

	s.Run("chan of chan", func() {
		reqs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		// 存放结果的channel的channel
		outs := make(chan chan int, len(reqs)) // order of result
		var wg sync.WaitGroup
		wg.Add(len(reqs))
		for _, x := range reqs {
			o := handle(&wg, x)
			outs <- o
		}

		go func() {
			wg.Wait()
			close(outs)
		}()

		// 读取结果，结果有序
		for o := range outs {
			fmt.Println(<-o)
		}
	})
}

// handle 处理请求，耗时随机模拟
func handle(wg *sync.WaitGroup, a int) chan int {
	out := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		out <- a * a
		wg.Done()
	}()
	return out
}

func doSomething(du time.Duration) <-chan int {
	outCh := make(chan int)
	go func() {
		<-time.After(du)
	}()
	return outCh
}

func TestChannelSuite(t *testing.T) {
	suite.Run(t, &ChannelSuite{})
}

type ChannelSuite struct {
	suite.Suite
}
