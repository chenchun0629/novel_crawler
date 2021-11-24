package serve

import (
	"context"
)

type Command struct {
}

func (c Command) Start(ctx context.Context) error {
	panic("implement me")
}

func (c Command) Stop() error {
	panic("implement me")
}

func (c Command) ForceStop() error {
	panic("implement me")
}
