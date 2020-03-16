package cronjob

import (
	"log"
	"os"
	"testing"

	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	i := 0
	spec := "* * * * * *"
	log.SetOutput(os.Stdout)
	log.Printf("cron wit spec: %s", spec)

	c := cron.New()
	c.AddFunc(spec, func() {
		i++
		log.Println(i)
	})
	c.Start()
	select {}
}
