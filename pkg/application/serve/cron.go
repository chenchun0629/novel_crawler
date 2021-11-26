package serve

import (
	"context"
	"github.com/novel_crawler/pkg/log"
	"github.com/robfig/cron/v3"
	"time"
)

type (
	Entry    = cron.Entry
	EntryID  = cron.EntryID
	Job      = cron.Job
	Schedule = cron.Schedule
)

type xcron struct {
	cl *cron.Cron

	isRunning bool
}

func NewXcron() *xcron {
	return &xcron{
		cl: cron.New(
			cron.WithLocation(time.Now().Location()),
		),
		isRunning: false,
	}
}

func (c *xcron) Remove(id EntryID) {
	c.cl.Remove(id)
}
func (c *xcron) Location() *time.Location {
	return c.cl.Location()
}

func (c *xcron) Entry(id EntryID) Entry {
	return c.cl.Entry(id)
}

func (c *xcron) Entries() []Entry {
	return c.cl.Entries()
}
func (c *xcron) Schedule(schedule Schedule, cmd Job) EntryID {
	return c.cl.Schedule(schedule, cmd)
}

func (c *xcron) AddJob(spec string, cmd Job) (EntryID, error) {
	return c.cl.AddJob(spec, cmd)
}

func (c *xcron) AddFunc(spec string, cmd func()) (EntryID, error) {
	return c.cl.AddFunc(spec, cmd)
}

func (c *xcron) Start(ctx context.Context) error {
	c.isRunning = true
	defer func() {
		c.isRunning = false
	}()

	go func() {
		c.cl.Start()
	}()

	<-ctx.Done()
	return nil
}

func (c *xcron) Stop() error {
	if !c.isRunning {
		return nil
	}

	log.Infow("stop cron serve ...")

	var ctx = c.cl.Stop()
	<-ctx.Done()

	log.Infow("stop cron serve success ...")
	return ctx.Err()
}

func (c *xcron) ForceStop() error {
	if !c.isRunning {
		return nil
	}

	log.Infow("force stop cron serve ...")

	return nil
}

func (c *xcron) IsRunning() bool {
	return c.isRunning
}
