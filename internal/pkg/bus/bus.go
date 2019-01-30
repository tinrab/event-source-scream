package bus

import (
	"fmt"

	"github.com/nats-io/go-nats"

	"github.com/tinrab/event-source-store/internal/pkg/config"
	"github.com/tinrab/event-source-store/internal/pkg/event"
)

type Bus struct {
	cfg  config.BusConfig
	conn *nats.Conn
}

type SubscribeHandler func(channel string, data []byte)

type Message struct {
	Data    []byte
	Channel string
}

func NewBus(c config.BusConfig) *Bus {
	return &Bus{
		cfg: c,
	}
}

func (b *Bus) Open() error {
	url := fmt.Sprintf("%s:%d", b.cfg.Host, b.cfg.Port)
	conn, err := nats.Connect(url)
	if err != nil {
		return err
	}
	b.conn = conn
	return nil
}

func (b *Bus) Close() {
	if b.conn != nil && !b.conn.IsClosed() {
		b.conn.Close()
	}
}

func (b *Bus) Publish(e event.Event) error {
	subj := fmt.Sprintf("%s:%d", e.Kind(), e.AggregateID())
	return b.conn.Publish(subj, e.Data())
}

func (b *Bus) Subscribe(channel string, handler SubscribeHandler) error {
	_, err := b.conn.Subscribe(channel, func(msg *nats.Msg) {
		handler(msg.Subject, msg.Data)
	})

	return err
}

func (b *Bus) SubscribeChan(channel string) (chan Message, error) {
	ch := make(chan *nats.Msg)
	res := make(chan Message)

	s, err := b.conn.ChanSubscribe(channel, ch)
	if err != nil {
		return nil, err
	}

	go func() {
		for m := range ch {
			res <- Message{
				Data:    m.Data,
				Channel: m.Subject,
			}
		}
		_ = s.Unsubscribe()
		close(res)
	}()

	return res, nil
}