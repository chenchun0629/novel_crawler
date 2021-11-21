package xgo

import "time"

func newEventBusOpts() *eventBusOpts {
	return &eventBusOpts{
		chanBuffer:     4,
		eventWorkers:   4,
		autoRetryTimes: 2,
		retryInterval:  time.Second,
		timeout:        5 * time.Second,
	}
}

type eventBusOpts struct {
	chanBuffer     int
	eventWorkers   int
	autoRetryTimes int
	retryInterval  time.Duration
	timeout        time.Duration
}

func (opts *eventBusOpts) apply(fns ...EventBusOptsFunc) {
	for _, fn := range fns {
		fn(opts)
	}
}

type EventBusOptsFunc func(opts *eventBusOpts)

func SetEventBusOptsChanBuffer(buffer int) EventBusOptsFunc {
	return func(opts *eventBusOpts) {
		opts.chanBuffer = buffer
	}
}

func SetEventBusOptsEventWorkers(workers int) EventBusOptsFunc {
	return func(opts *eventBusOpts) {
		opts.eventWorkers = workers
	}
}

func SetEventBusOptsAutoRetryTimes(times int) EventBusOptsFunc {
	return func(opts *eventBusOpts) {
		opts.autoRetryTimes = times
	}
}

func SetEventBusOptsRetryInterval(retryInterval time.Duration) EventBusOptsFunc {
	return func(opts *eventBusOpts) {
		opts.retryInterval = retryInterval
	}
}

func SetEventBusOptsTimeout(timeout time.Duration) EventBusOptsFunc {
	return func(opts *eventBusOpts) {
		opts.timeout = timeout
	}
}
