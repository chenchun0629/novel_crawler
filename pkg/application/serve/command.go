package serve

import (
	"context"
)

type Command struct {
	fn func() error

	isRunning bool
}

func (c *Command) IsRunning() bool {
	return c.isRunning
}

func NewCommand(fn func() error) *Command {
	return &Command{fn: fn}
}

func (c *Command) Start(ctx context.Context) error {
	c.isRunning = true
	defer func() {
		c.isRunning = false
	}()

	return c.fn()
}

func (c Command) Stop() error {
	if !c.isRunning {
		return nil
	}

	return nil
}

func (c Command) ForceStop() error {
	if !c.isRunning {
		return nil
	}

	return nil
}
