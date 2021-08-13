package messager

import (
	"context"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	nc NatsConnection
}

func NewNatsPublisher(nc NatsConnectionCloser) *NatsPublisher {
	return &NatsPublisher{
		nc: nc,
	}
}

func (p NatsPublisher) Publish(ctx context.Context, subject, key string, header MessageHeaders, messsage []byte) error {
	head := nats.Header{}
	for k, h := range header {
		head.Add(k, h)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    messsage,
		Header:  head,
	}

	return p.nc.Publish(ctx, msg)
}

func (p NatsPublisher) Request(ctx context.Context, subject string, header MessageHeaders, messsage []byte) (m *nats.Msg, err error) {
	head := nats.Header{}
	for k, h := range header {
		head.Add(k, h)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    messsage,
		Header:  head,
	}

	return p.nc.Request(ctx, msg)
}
