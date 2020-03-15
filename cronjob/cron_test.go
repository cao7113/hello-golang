package cronjob

import (
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func TestCron(t *testing.T) {
	i := 0
	spec := "* * * * * *"
	logrus.Infof("cron wit spec: %s", spec)

	c := cron.New()
	c.AddFunc(spec, func() {
		i++
		//log.Println("cron running:", i)
		logrus.Infoln(i)
	})
	c.Start()
	select {}
}
