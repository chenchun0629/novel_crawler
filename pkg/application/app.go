package application

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	BeforeAppStart = iota
	AfterAppStart
)

type App struct {
	ctx     context.Context
	appInfo appInfo

	callbacks map[int][]func(ctx context.Context) error // 生命周期

	serves []Serve
	sigs   []os.Signal
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx:     ctx,
		appInfo: appInfo{},
		serves:  make([]Serve, 0),
		sigs:    []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
}

type appInfo struct {
	id       string
	name     string
	version  string
	metadata map[string]string
}

func (a *App) RegisterStartup(fn ...func(ctx context.Context) error) {
	a.callbacks[BeforeAppStart] = append(a.callbacks[BeforeAppStart], fn...)
}

func (a *App) Run() error {
	var (
		eg, ctx = errgroup.WithContext(a.ctx)
		wg      = sync.WaitGroup{}
	)

	if fns, has := a.callbacks[BeforeAppStart]; has {
		for _, fn := range fns {
			if err := fn(ctx); err != nil {
				return err
			}
		}
	}

	for _, serv := range a.serves {
		var serv = serv
		wg.Add(1)
		eg.Go(func() error {
			return serv.Start(ctx)
		})
	}

	wg.Wait()

	if fns, has := a.callbacks[AfterAppStart]; has {
		for _, fn := range fns {
			if err := fn(ctx); err != nil {
				return err
			}
		}
	}

	// 注册 和 处理 signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				err := a.Stop()
				if err != nil {
					//a.opts.logger.Errorf("failed to stop app: %v", err)
					return err
				}
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a App) Stop() error {
	var eg = errgroup.Group{}
	for _, serv := range a.serves {
		var serv = serv
		eg.Go(func() error {
			return serv.Stop()
		})
	}

	return eg.Wait()
}

func (a App) ForceStop() error {
	var eg = errgroup.Group{}
	for _, serv := range a.serves {
		var serv = serv
		eg.Go(func() error {
			return serv.ForceStop()
		})
	}

	return eg.Wait()
}
