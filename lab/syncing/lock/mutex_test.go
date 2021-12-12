package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 互斥锁针对读写不确定的情景，解决并发中的资源共享时的同步问题，属于较低层API，顶层优先使用channel解决
func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 3; r++ {
		go func() {
			mutex.Lock()
			fmt.Printf("before adding count=%d \n", count)
			count++
			fmt.Printf("afeter adding count=%d \n", count)
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)
}

/*
output:

before adding count=0
afeter adding count=1
before adding count=1
afeter adding count=2
before adding count=2
afeter adding count=3
the count is :  3
*/

// 读写锁针对 读多写少的场景，提升互斥锁的效率
func TestRwMutex(t *testing.T) {
	var mutex sync.RWMutex
	arr := []int{1, 2, 3}

	go func() {
		fmt.Println("Try to lock reading operation.")
		mutex.RLock()
		fmt.Println("The reading operation is locked.")

		fmt.Println("The len of arr is : ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	go func() {
		fmt.Println("Try to lock writing operation.")
		mutex.Lock()
		fmt.Println("Writing operation is locked.")

		arr = append(arr, 4)

		fmt.Println("Try to unlock writing operation.")
		mutex.Unlock()
		fmt.Println("Writing operation is unlocked.")
	}()

	go func() {
		fmt.Println("Try to lock reading operation.")
		mutex.RLock()
		fmt.Println("The reading operation is locked.")

		fmt.Println("The len of arr is : ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	time.Sleep(time.Second * 2)
}

func TestWaitGroup(t *testing.T) {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait()
}
