package messager

import (
	"context"

	"github.com/nats-io/nats.go"
)

type NatsConnection interface {
	Publish(ctx context.Context, msg *nats.Msg) (e error)
	Subscribe(subject string, msg nats.MsgHandler) (s *nats.Subscription, e error)
	Request(ctx context.Context, msg *nats.Msg) (m *nats.Msg, e error)
}

type NatsConnectionCloser interface {
	NatsConnection
	Closer
}

type NatsRequester interface {
	Request(ctx context.Context, subject string, header MessageHeaders, messsage []byte) (*nats.Msg, error)
}

type NatsReplyRequester interface {
	Reply(ctx context.Context, msg interface{}) (resp []byte, e error)
}

type NatsPubReq interface {
	Publisher
	NatsRequester
}

type NatsReplyEventHandler interface {
	NatsReplyRequester
	EventHandler
}
