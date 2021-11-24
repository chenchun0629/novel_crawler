package serve

import (
	"context"
)

type Command struct {
	fn func() error
}

func NewCommand(fn func() error) *Command {
	return &Command{fn: fn}
}

func (c Command) Start(ctx context.Context) error {
	return c.fn()
}

func (c Command) Stop() error {
	return nil
}

func (c Command) ForceStop() error {
	return nil
}
