package xgo

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	var cnt = 0
	var task = func(items []interface{}) error {
		time.Sleep(0 * time.Second)
		fmt.Println(items)
		cnt += len(items)
		return nil
	}
	var agg = NewAggregator(task, SetAggregatorOptionBuffer(10))
	agg.Run()

	defaultManager.print()

	for i := 0; i < 100; i++ {
		agg.Push(i)
	}

	agg.Close()
	defaultManager.print()
	fmt.Println(cnt)
}
