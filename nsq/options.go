package nsq

import (
	"context"

	"github.com/rai-project/broker"
)

const (
	maxInFlightKey            = "github.com/rai-project/broker/nsq/maxInFlight"
	concurrentHandlerCountKey = "github.com/rai-project/broker/nsq/concurrentHandlerCount"
)

var (
	DefaultConcurrentHandlerCount = 1
)

func MaxInFlight(n int) broker.Option {
	return func(o *broker.Options) {
		o.Context = context.WithValue(o.Context, maxInFlightKey, n)
	}
}

func ConcurrentHandlerCount(n int) broker.SubscribeOption {
	return func(o *broker.SubscribeOptions) {
		o.Context = context.WithValue(o.Context, concurrentHandlerCountKey, n)
	}
}
