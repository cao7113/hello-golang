package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("start", i)
	})
	c.Start()
	select {} // 阻塞主线程不退出
}
