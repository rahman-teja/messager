package messager

import (
	"context"
)

type MessageHeaders map[string]string

type Closer interface {
	Close() (e error)
}

type DLQHandler interface {
	Publish(ctx context.Context, msg *DeadLetterQueueMessage) error
}

type Publisher interface {
	Publish(ctx context.Context, subject, key string, header MessageHeaders, messsage []byte) error
}

type EventHandler interface {
	Handle(ctx context.Context, msg interface{}) error
}

type Subscriber interface {
	Subscribe() error
}

type SubscriberCloser interface {
	Subscriber
	Closer
}
