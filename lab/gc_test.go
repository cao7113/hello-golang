package lab

import (
	"log"
	"runtime"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Close() {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Close")
}

func (p *Person) NewOpen() {
	log.Println("init")
	runtime.SetFinalizer(p, (*Person).Close)
}

func Tt(p *Person) {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Tt")
}

func Mem(m *runtime.MemStats) {
	runtime.ReadMemStats(m)
	log.Printf("%d Kb\n", m.Alloc/1024)
}

// p122 《深入学习go语言》
func (s *LabSuite) TestGc() {
	var m runtime.MemStats
	Mem(&m)

	p := &Person{Name: "lee", Age: 4}
	p.NewOpen()
	log.Println("Gc 完成第一次")
	log.Println("p: ", p)
	// 显示调用GC()可以显示触发垃圾回收
	runtime.GC()
	time.Sleep(5 * time.Second)
	Mem(&m)

	p1 := &Person{Name: "Goo", Age: 9}
	runtime.SetFinalizer(p1, Tt)
	log.Println("Gc 完成第二次")
	time.Sleep(2 * time.Second)
	runtime.GC()
	time.Sleep(2 * time.Second)
	Mem(&m)
}

/*
环境变量GOGC 设置垃圾回收百分比，默认 100
当新分配的内存与上次垃圾回收后剩余的实时数据的比率达到此阈值后触发垃圾回收
*/
