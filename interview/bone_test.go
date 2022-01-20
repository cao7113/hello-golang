package interview

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"runtime"
	"sync"
	"testing"
)

func (s *InterviewSuite) TestChannel() {
	s.Run("should panic on: send on closed channel", func() {
		s.PanicsWithError("send on closed channel", func() {
			c := make(chan string, 10)
			close(c)
			c <- "test string"
		})
	})

	s.Run("read closed channel return zero value", func() {
		c := make(chan string, 10)
		close(c)
		//for {
		//	select {
		//	case s := <-c: // always hit here
		//		fmt.Printf("received: %s \n", s)
		//	default:
		//		fmt.Println("no data")
		//	}
		//}
		cc := <-c
		s.Equal("", cc)
		cc = <-c
		s.Equal("", cc)
	})

	s.Run("for range closed channel", func() {
		c := make(chan string, 10)
		close(c)
		for s := range c { // not run here
			fmt.Printf("received: %s \n", s)
		}
	})
}

/*
type Saying interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestAssign(t *testing.T) {
	// Student not impl Saying method, compile error
	//var stu Saying = Student{}
}
*/

// Q3
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func (s *InterviewSuite) TestShowA() {
	t := Teacher{}
	t.ShowA()
}

func (s *InterviewSuite) TestWg() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)

			wg.Done()
		}(i)
	}
	wg.Wait()
}

type student struct {
	Name string
	Age  int
}

func (s *InterviewSuite) TestForVarAddr() {
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	m := make(map[string]*student)
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for _, v := range m {
		s.Equal("wang", v.Name)
	}
}

func (s *InterviewSuite) TestDeferOrder() {
	s.PanicsWithValue("panic!", deferCall)
}

func deferCall() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()

	panic("panic!")
	fmt.Printf("hello defer") // never reach here
}

func TestInterviewSuite(t *testing.T) {
	suite.Run(t, &InterviewSuite{})
}

type InterviewSuite struct {
	suite.Suite
}
