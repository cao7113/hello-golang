package lab

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func (s *ChannelSuite) TestWorkerModel() {
	t0 := time.Now()
	jobCnt := 1_000_000
	pool := genJob(jobCnt)
	resultCh := make(chan int, 3)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			wn := fmt.Sprintf("worker%d", i+1)
			workerDo(wn, pool, resultCh)
			wg.Done()
		}(i)
	}

	s.Run("wait collect routine", func() {
		wg.Add(1)
		go func() { // collect result
			total := 0
			cnt := 0
			for i := range resultCh {
				//log.Println(cnt, "got result: ", i)
				total += i
				cnt++
				if jobCnt == cnt { // 判断 全部工作结束
					close(resultCh)
				}
			}
			log.Println("== final result: ", total)
			wg.Done()
		}()

		wg.Wait()
		du := time.Until(t0)
		log.Println("taken ms: ", du.Milliseconds())
	})

	//s.Run("wait workers only", func() {
	//	wg.Wait()
	//
	//	total := 0
	//	cnt := 0
	//	for i := range resultCh {
	//		//log.Println(cnt, "got result: ", i)
	//		total += i
	//		cnt++
	//		if jobCnt == cnt { // 判断 全部工作结束
	//			close(resultCh)
	//		}
	//	}
	//	log.Println("== final result: ", total)
	//})
}

func workerDo(name string, inCh <-chan int, outCh chan<- int) {
	log.Println(name, "starting")
	for j := range inCh {
		//log.Println(name, "handling job", j)
		outCh <- j * j
	}
	log.Println(name, "has done")
}

func genJob(cnt int) <-chan int {
	pool := make(chan int, 3)
	go func() {
		for i := 0; i < cnt; i++ {
			//log.Println("gen job", i)
			pool <- i + 1
		}
		close(pool)
		log.Println("gen", cnt, "jobs done and pool-channel closed")
	}()
	return pool
}
