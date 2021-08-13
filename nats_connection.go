package messager

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsConn struct {
	timeoutInScd int
	service      string
	nc           *nats.Conn
}

func NewNatsConn(timeoutInScd int, service, url string, options ...nats.Option) (n *NatsConn, e error) {
	var nc *nats.Conn

	nc, e = nats.Connect(url, options...)
	if e != nil {
		return
	}

	n = &NatsConn{
		nc:           nc,
		timeoutInScd: timeoutInScd,
		service:      service,
	}
	return
}

func (c NatsConn) Publish(ctx context.Context, msg *nats.Msg) (e error) {
	e = c.nc.PublishMsg(msg)
	return
}

func (c NatsConn) Subscribe(subject string, msg nats.MsgHandler) (s *nats.Subscription, e error) {
	return c.nc.Subscribe(subject, msg)
}

func (c NatsConn) Request(ctx context.Context, msg *nats.Msg) (m *nats.Msg, e error) {
	cx, cc := context.WithTimeout(ctx, time.Second*time.Duration(c.timeoutInScd))
	defer cc()

	return c.nc.RequestMsgWithContext(cx, msg)
}

func (c NatsConn) Close() (e error) {
	e = c.nc.Flush()
	if e != nil {
		return e
	}

	c.nc.Close()

	return nil
}
