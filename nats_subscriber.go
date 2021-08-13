package messager

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsSubscriber struct {
	nc          NatsConnection
	serviceName string
	topic       string
	handler     EventHandler
	dlqHandler  DLQHandler
}

func NewNatsSubscriber(nc NatsConnectionCloser, service, topic string, handler EventHandler, dlq DLQHandler) *NatsSubscriber {
	return &NatsSubscriber{
		nc:          nc,
		serviceName: service,
		topic:       topic,
		handler:     handler,
		dlqHandler:  dlq,
	}
}

func (n NatsSubscriber) Subscribe() (e error) {
	n.nc.Subscribe(n.topic, func(msg *nats.Msg) {
		ctx := context.Background()

		e = n.handler.Handle(ctx, msg)
		if e != nil {
			n.sendToDLQ(ctx, msg, e)
		}
	})

	return nil
}

func (n NatsSubscriber) sendToDLQ(context context.Context, message *nats.Msg, err error) {
	if n.dlqHandler == nil {
		return
	}

	dlqMessage := DeadLetterQueueMessage{
		Subject:           message.Subject,
		Consumer:          n.serviceName,
		FailedConsumeDate: time.Now().Format(time.RFC3339Nano),
		Message:           message.Data,
		CausedBy:          err.Error(),
	}

	n.dlqHandler.Publish(context, &dlqMessage)
}

func (n NatsSubscriber) Close() error {
	return nil
}
