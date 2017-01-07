package schedule

import (
	"github.com/labstack/gommon/log"
	"time"
)

type recurrences interface {
	Every()
}

type constraints interface {
	Weekdays()
	Workdays()
}

type hoocks interface {
	After()
	Before()
	Timezone()
}

type additionals interface {
	Between()
	When()
	Cron()
	On()
}

type job struct {
	command func()
}

type cron struct {
	jobs []job
	utc  *time.Location
	Minute
	Hour
	Day
	Months
}

func (c *cron) Timezone(zone string) cron {
	utc, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}
	c.utc = utc
	return c
}

func Command(f func()) cron {
	c := cron{}
	c.jobs = append(c.jobs, job{command: f})
	return c
}
