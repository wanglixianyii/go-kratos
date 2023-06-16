package redis

import (
	"errors"
	"strings"
	"time"
	broker2 "user-api-job/pkg/broker"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
)

const (
	defaultBroker = "redis://127.0.0.1:6379"
)

type redisBroker struct {
	addr       string
	pool       *redis.Pool
	opts       broker2.Options
	commonOpts *commonOptions
}

// NewBroker returns a new common implemented using the Redis pub/sub
// protocol. The connection address may be a fully qualified IANA address such
// as: redis://user-api:secret@localhost:6379/0?foo=bar&qux=baz
func NewBroker(opts ...broker2.Option) broker2.Broker {
	commonOpts := &commonOptions{
		maxIdle:        DefaultMaxIdle,
		maxActive:      DefaultMaxActive,
		idleTimeout:    DefaultIdleTimeout,
		connectTimeout: DefaultConnectTimeout,
		readTimeout:    DefaultReadTimeout,
		writeTimeout:   DefaultWriteTimeout,
	}

	options := broker2.NewOptionsAndApply(opts...)

	return &redisBroker{
		opts:       options,
		commonOpts: commonOpts,
	}
}

func (b *redisBroker) Name() string {
	return "redis"
}

func (b *redisBroker) Options() broker2.Options {
	return b.opts
}

func (b *redisBroker) Address() string {
	return b.addr
}

func (b *redisBroker) Init(opts ...broker2.Option) error {
	if b.pool != nil {
		return errors.New("redis: cannot init while connected")
	}

	var addr string

	if len(b.opts.Addrs) == 0 || b.opts.Addrs[0] == "" {
		addr = defaultBroker
	} else {
		addr = b.opts.Addrs[0]

		if !strings.HasPrefix(addr, "redis://") {
			addr = "redis://" + addr
		}
	}

	b.addr = addr

	b.opts.Apply(opts...)

	if v, ok := b.opts.Context.Value(optionsKey).(*commonOptions); ok {
		b.commonOpts = v
	}

	return nil
}

func (b *redisBroker) Connect() error {
	if b.pool != nil {
		return nil
	}

	b.pool = &redis.Pool{
		MaxIdle:     b.commonOpts.maxIdle,
		MaxActive:   b.commonOpts.maxActive,
		IdleTimeout: b.commonOpts.idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(
				b.addr,
				redis.DialConnectTimeout(b.commonOpts.connectTimeout),
				redis.DialReadTimeout(DefaultHealthCheckPeriod+b.commonOpts.readTimeout),
				redis.DialWriteTimeout(b.commonOpts.writeTimeout),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if nil != err {
				log.Error("[redis] ping error:" + err.Error())
			}
			return err
		},
	}

	return nil
}

func (b *redisBroker) Disconnect() error {
	err := b.pool.Close()
	b.pool = nil
	b.addr = ""
	return err
}

func (b *redisBroker) Publish(topic string, msg broker2.Any, opts ...broker2.PublishOption) error {
	buf, err := broker2.Marshal(b.opts.Codec, msg)
	if err != nil {
		return err
	}

	return b.publish(topic, buf, opts...)
}

func (b *redisBroker) publish(topic string, msg []byte, _ ...broker2.PublishOption) error {
	conn := b.pool.Get()
	_, err := redis.Int(conn.Do("PUBLISH", topic, msg))
	_ = conn.Close()
	return err
}

func (b *redisBroker) Subscribe(topic string, handler broker2.Handler, binder broker2.Binder, opts ...broker2.SubscribeOption) (broker2.Subscriber, error) {
	var options broker2.SubscribeOptions
	for _, o := range opts {
		o(&options)
	}

	s := subscriber{
		codec:   b.opts.Codec,
		conn:    &redis.PubSubConn{Conn: b.pool.Get()},
		topic:   topic,
		handler: handler,
		binder:  binder,
		opts:    options,
	}

	if err := s.conn.Subscribe(s.topic); err != nil {
		return nil, err
	}

	go s.recv()

	return &s, nil
}
