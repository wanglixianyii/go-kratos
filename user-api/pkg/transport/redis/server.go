package redis

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"net/url"
	"strings"
	"sync"
	"time"
	broker2 "user-api-job/pkg/broker"
	"user-job/pkg/broker/redis"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type SubscriberMap map[string]broker2.Subscriber

type SubscribeOption struct {
	handler broker2.Handler
	binder  broker2.Binder
	opts    []broker2.SubscribeOption
}
type SubscribeOptionMap map[string]*SubscribeOption

type Server struct {
	broker2.Broker
	brokerOpts []broker2.Option

	subscribers    SubscriberMap
	subscriberOpts SubscribeOptionMap

	sync.RWMutex
	started bool

	baseCtx context.Context
	err     error
}

func NewServer(opts ...ServerOption) *Server {
	opts = append(opts, WithReadTimeout(24*time.Hour))
	opts = append(opts, WithIdleTimeout(24*time.Hour))

	srv := &Server{
		baseCtx:        context.Background(),
		subscribers:    SubscriberMap{},
		subscriberOpts: SubscribeOptionMap{},
		brokerOpts:     []broker2.Option{},
		started:        false,
	}

	srv.init(opts...)

	srv.Broker = redis.NewBroker(srv.brokerOpts...)

	return srv
}

func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}

func (s *Server) Name() string {
	return "redis"
}

func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}

	addr := s.Address()
	if !strings.HasPrefix(addr, "tcp://") {
		addr = "tcp://" + addr
	}

	return url.Parse(addr)
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}

	if s.started {
		return nil
	}

	s.err = s.Init()
	if s.err != nil {
		log.Errorf("[redis] init broker failed: [%s]", s.err.Error())
		return s.err
	}

	s.err = s.Connect()
	if s.err != nil {
		return s.err
	}

	log.Infof("[redis] server listening on: %s", s.Address())

	s.err = s.doRegisterSubscriberMap()
	if s.err != nil {
		return s.err
	}

	s.baseCtx = ctx
	s.started = true

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	log.Info("[redis] server stopping")
	s.started = false
	return s.Disconnect()
}

func (s *Server) RegisterSubscriber(topic string, handler broker2.Handler, binder broker2.Binder, opts ...broker2.SubscribeOption) error {
	s.Lock()
	defer s.Unlock()

	if s.started {
		return s.doRegisterSubscriber(topic, handler, binder, opts...)
	} else {
		s.subscriberOpts[topic] = &SubscribeOption{handler: handler, binder: binder, opts: opts}
	}
	return nil
}

func (s *Server) doRegisterSubscriber(topic string, handler broker2.Handler, binder broker2.Binder, opts ...broker2.SubscribeOption) error {
	sub, err := s.Subscribe(topic, handler, binder, opts...)
	if err != nil {
		return err
	}

	s.subscribers[topic] = sub

	return nil
}

func (s *Server) doRegisterSubscriberMap() error {
	for topic, opt := range s.subscriberOpts {
		_ = s.doRegisterSubscriber(topic, opt.handler, opt.binder, opt.opts...)
	}
	s.subscriberOpts = SubscribeOptionMap{}
	return nil
}
