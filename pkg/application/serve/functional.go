package serve

import (
	"context"
)

type Functional struct {
	fn func() error

	isRunning bool
}

func (c *Functional) IsRunning() bool {
	return c.isRunning
}

func NewFunctional(fn func() error) *Functional {
	return &Functional{fn: fn}
}

func (c *Functional) Start(ctx context.Context) error {
	c.isRunning = true
	defer func() {
		c.isRunning = false
	}()

	return c.fn()
}

func (c Functional) Stop() error {
	if !c.isRunning {
		return nil
	}

	return nil
}

func (c Functional) ForceStop() error {
	if !c.isRunning {
		return nil
	}

	return nil
}
