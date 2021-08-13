package messager

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type DeadLetterQueueMessage struct {
	Subject           string         `json:"subject"`
	Publisher         string         `json:"publisher"`
	Consumer          string         `json:"consumer"`
	Key               string         `json:"key"`
	Headers           MessageHeaders `json:"headers"`
	Message           []byte         `json:"message"`
	CausedBy          string         `json:"causedBy"`
	FailedConsumeDate string         `json:"failedConsumeDate"`
}

type DLQHandlerAdapter struct {
	topic     string
	publisher Publisher
}

func NewDLQHandlerAdapter(topic string, publisher Publisher) *DLQHandlerAdapter {
	return &DLQHandlerAdapter{
		topic:     topic,
		publisher: publisher,
	}
}

func (d *DLQHandlerAdapter) Publish(ctx context.Context, msg *DeadLetterQueueMessage) (err error) {
	if d == nil {
		return ErrInvalidPublsiher
	}

	key := fmt.Sprintf("%s:%s:%s:%d",
		msg.Consumer,
		msg.Subject,
		msg.Key,
		time.Now().UnixNano(),
	)

	messageByte, _ := json.Marshal(msg)

	err = d.publisher.Publish(ctx, d.topic, key, nil, messageByte)
	return
}
