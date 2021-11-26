package application

import (
	"context"
	"errors"
	"github.com/novel_crawler/pkg/log"
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
	ctx    context.Context
	cancel func()
	lock   sync.Mutex

	appInfo appInfo

	callbacks map[int][]func(ctx context.Context) error // 生命周期

	serves []Serve
	sigs   []os.Signal
}

func NewApp(ctx context.Context) *App {
	var c, cancel = context.WithCancel(ctx)
	var a = &App{
		ctx:    c,
		cancel: cancel,
	}

	return a.init()
}

type appInfo struct {
	id       string
	name     string
	version  string
	metadata map[string]string
}

func (a *App) init() *App {

	a.appInfo = appInfo{
		id:       "",
		name:     "novel_crawler",
		version:  "",
		metadata: nil,
	}
	a.serves = make([]Serve, 0)
	a.sigs = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	a.callbacks = make(map[int][]func(ctx context.Context) error)

	for _, v := range []int{BeforeAppStart, AfterAppStart} {
		a.callbacks[v] = make([]func(ctx context.Context) error, 0)
	}

	return a
}

func (a *App) RegisterServes(serves ...Serve) {
	a.serves = append(a.serves, serves...)
}

func (a *App) RegisterBeforeAppStart(fn ...func(ctx context.Context) error) {
	a.callbacks[BeforeAppStart] = append(a.callbacks[BeforeAppStart], fn...)
}
func (a *App) RegisterAfterAppStart(fn ...func(ctx context.Context) error) {
	a.callbacks[AfterAppStart] = append(a.callbacks[AfterAppStart], fn...)
}

func (a *App) Run() error {
	if len(a.serves) == 0 {
		return errors.New("empty serve, application cannot run")
	}

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
			defer func() {
				a.tryCancel()
			}()
			wg.Done()
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
	var c = make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				log.Infow("signal stop application ...")
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

func (a *App) tryCancel() {
	a.lock.Lock()
	defer a.lock.Unlock()

	for _, serv := range a.serves {
		if serv.IsRunning() {
			return
		}
	}

	a.cancel()
}

func (a *App) Stop() error {
	defer func() {
		a.cancel()
	}()

	log.Infow("start to stop application ...")

	var eg = errgroup.Group{}
	for _, serv := range a.serves {
		var serv = serv
		eg.Go(func() error {
			return serv.Stop()
		})
	}

	return eg.Wait()
}

func (a *App) ForceStop() error {
	defer func() {
		a.cancel()
	}()

	log.Infow("start to force stop application ...")

	var eg = errgroup.Group{}
	for _, serv := range a.serves {
		var serv = serv
		eg.Go(func() error {
			return serv.ForceStop()
		})
	}

	return eg.Wait()
}
