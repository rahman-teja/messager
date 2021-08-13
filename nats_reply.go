package messager

import (
	"context"

	"github.com/nats-io/nats.go"
)

type NatsReply struct {
	nc      NatsConnection
	service string
	topic   string
	handler NatsReplyRequester
}

func NewNatsReply(nc NatsConnectionCloser, service, topic string, handler NatsReplyRequester) *NatsReply {
	return &NatsReply{
		nc:      nc,
		service: service,
		topic:   topic,
		handler: handler,
	}
}

func (n NatsReply) Subscribe() (e error) {
	n.nc.Subscribe(n.topic, func(msg *nats.Msg) {
		var bts []byte

		bts, _ = n.handler.Reply(context.Background(), msg)

		msg.Respond(bts)
	})

	return nil
}

func (n NatsReply) Close() error {
	return nil
}
