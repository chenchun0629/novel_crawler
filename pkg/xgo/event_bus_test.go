package xgo

import (
	"context"
	"fmt"
	"github.com/Ksloveyuan/channelx"
	"testing"
	"time"
)

type TEBEvent struct {
	Name string
}

func NewTEBEvent(name string) *TEBEvent {
	return &TEBEvent{Name: name}
}

func (T TEBEvent) ID() EventID {
	return 1
}

type TEBEventHandler struct{}

func (T TEBEventHandler) OnEvent(ctx context.Context, event channelx.Event) error {
	fmt.Println("============hello, this is on event", event)
	return nil
}

func (T TEBEventHandler) Logger() channelx.Logger {
	// todo logger helper
	return channelx.NewConsoleLogger()
}

func (T TEBEventHandler) CanAutoRetry(err error) bool {
	return true
}

func TestEventBus(t *testing.T) {
	var eb = NewEventBus()
	eb.Subscribe(1, &TEBEventHandler{})

	eb.Run()

	eb.Publish(NewTEBEvent("aa"))
	eb.Publish(NewTEBEvent("bb"))
	eb.Publish(NewTEBEvent("cc"))
	eb.Publish(NewTEBEvent("dd"))

	time.Sleep(time.Hour)
}
