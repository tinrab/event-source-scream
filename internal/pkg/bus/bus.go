package bus

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"

	"github.com/tinrab/event-source-scream/internal/pkg/command"
	"github.com/tinrab/event-source-scream/internal/pkg/config"
	"github.com/tinrab/event-source-scream/internal/pkg/event"
)

type Bus struct {
	cfg  config.BusConfig
	conn *nats.Conn
}

type EventHandler func(event.Event)

type CommandHandler func(kind string, data []byte) command.Result

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

func (b *Bus) PublishEvent(e event.Event) error {
	subj := fmt.Sprintf("%s:%d", e.Kind, e.AggregateID)
	data, _ := json.Marshal(e)
	return b.conn.Publish(subj, data)
}

func (b *Bus) HandleEvent(kind string, handler EventHandler) error {
	_, err := b.conn.Subscribe(kind, func(msg *nats.Msg) {
		var e event.Event
		if err := json.Unmarshal(msg.Data, &e); err != nil {
			log.Print(err)
			return
		}

		e.Kind = msg.Subject

		handler(e)
	})

	return err
}

func (b *Bus) PublishCommand(c command.Command, res *command.Result, timeout time.Duration) error {
	data, _ := json.Marshal(c)
	msg, err := b.conn.Request(c.Kind, data, timeout)
	if err != nil {
		return err
	}
	return json.Unmarshal(msg.Data, res)
}

func (b *Bus) HandleCommand(kind string, handler CommandHandler) error {
	_, err := b.conn.Subscribe(kind, func(msg *nats.Msg) {
		//var c command.Command
		//if err := json.Unmarshal(msg.Data, &c); err != nil {
		//	log.Print(err)
		//	return
		//}

		//c.Kind = msg.Subject
		r := handler(msg.Subject, msg.Data)

		data, err := json.Marshal(r)
		if err != nil {
			log.Print(err)
			return
		}

		if err = b.conn.Publish(msg.Reply, data); err != nil {
			log.Print(err)
		}
	})
	return err
}
