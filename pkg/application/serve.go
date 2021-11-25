package application

import "context"

type Serve interface {
	Start(ctx context.Context) error
	Stop() error
	ForceStop() error
	IsRunning() bool
}
