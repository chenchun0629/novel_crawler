package xgo

import "github.com/Ksloveyuan/channelx"

type (
	Event        = channelx.Event
	EventID      = channelx.EventID
	JobStatus    = channelx.JobStatus
	EventHandler = channelx.EventHandler
)

func NewEventBus(optsFuncs ...EventBusOptsFunc) *eventBus {
	var opts = newEventBusOpts()
	opts.apply(optsFuncs...)
	return &eventBus{
		cl: channelx.NewEventBus(
			channelx.NewConsoleLogger(),
			opts.chanBuffer,
			opts.eventWorkers,
			opts.autoRetryTimes,
			opts.retryInterval,
			opts.timeout,
		),
		managerID: "",
	}
}

type eventBus struct {
	cl *channelx.EventBus

	managerID string
}

func (w *eventBus) Close() {
	w.cl.Stop()
	defaultManager.remove(w.managerID)
}

func (w *eventBus) Run() {
	w.cl.Start()
	w.managerID = defaultManager.register(w)
}

func (w *eventBus) Publish(item Event) <-chan JobStatus {
	return w.cl.Publish(item)
}

func (w *eventBus) Subscribe(eventID EventID, handlers ...EventHandler) {
	w.cl.Subscribe(eventID, handlers...)
}

func (w *eventBus) Unsubscribe(eventID EventID, handlers ...EventHandler) {
	w.cl.Unsubscribe(eventID, handlers...)
}
