package redis

import (
	"errors"
	"time"
	broker2 "user-rpc-rpc-api-user-rpc-job/pkg/broker"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
)

type subscriber struct {
	codec   encoding.Codec
	conn    *redis.PubSubConn
	topic   string
	handler broker2.Handler
	binder  broker2.Binder
	opts    broker2.SubscribeOptions
	done    chan error
}

func (s *subscriber) onStart() error {
	return nil
}

func (s *subscriber) onMessage(channel string, data []byte) error {
	var m broker2.Message

	if s.binder != nil {
		m.Body = s.binder()
	} else {
		m.Body = data
	}

	p := publication{
		topic:   channel,
		message: &m,
	}

	if p.err = broker2.Unmarshal(s.codec, data, &m.Body); p.err != nil {
		//log.Error("[redis]", err)
		return p.err
	}

	if p.err = s.handler(s.opts.Context, &p); p.err != nil {
		return p.err
	}

	if s.opts.AutoAck {
		if p.err = p.Ack(); p.err != nil {
			return p.err
		}
	}

	return nil
}

func (s *subscriber) ping() error {
	if s.conn == nil {
		return errors.New("cannot ping")
	}

	if err := s.conn.Ping(""); err != nil {
		return err
	}
	return nil
}

func (s *subscriber) recv() {
	defer func(conn *redis.PubSubConn) {
		err := conn.Close()
		if err != nil {
			log.Error("[redis] close pubsub connection error: ", err)
		}
	}(s.conn)

	s.done = make(chan error, 1)

	ticker := time.NewTicker(DefaultHealthCheckPeriod)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := s.ping(); err != nil {
					s.done <- err
					return
				}
			case <-s.opts.Context.Done():
				s.done <- nil
				return
			}
		}
	}()

	_ = s.ping()

	for {
		switch x := s.conn.Receive().(type) {
		case error:
			log.Errorf("[redis] recv error: %s\n", x.Error())
			s.done <- x
			return

		case redis.Message:
			if err := s.onMessage(x.Channel, x.Data); err != nil {
				s.done <- err
				break
			}

		case redis.Subscription:
			switch x.Count {
			case 0:
				s.done <- nil
				return
			}

		case redis.Pong:
			log.Debug("[redis] pong")
		}
	}
}

func (s *subscriber) Options() broker2.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Topic() string {
	return s.topic
}

func (s *subscriber) Unsubscribe() error {
	return s.conn.Unsubscribe()
}
