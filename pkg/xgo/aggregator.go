package xgo

import "github.com/Ksloveyuan/channelx"

type (
	Task                 = channelx.BatchProcessFunc
	AggregatorOptionFunc = channelx.SetAggregatorOptionFunc
)

func SetAggregatorOptionBuffer(buffer int) AggregatorOptionFunc {
	return func(option channelx.AggregatorOption) channelx.AggregatorOption {
		option.ChannelBufferSize = buffer
		return option
	}
}

func NewAggregator(task Task, optionFuncs ...AggregatorOptionFunc) *aggregator {
	var w = &aggregator{
		cl: channelx.NewAggregator(
			task,
			optionFuncs...,
		),
	}

	return w
}

type aggregator struct {
	cl *channelx.Aggregator

	managerID string
}

func (w *aggregator) Close() {
	w.cl.SafeStop()
	defaultManager.remove(w.managerID)
}

func (w *aggregator) Run() {
	w.cl.Start()
	w.managerID = defaultManager.register(w)
}

func (w *aggregator) Push(item interface{}) {
	w.cl.Enqueue(item)
}
