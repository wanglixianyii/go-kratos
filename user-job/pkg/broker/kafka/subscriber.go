package kafka

import (
	"sync"
	broker2 "user-rpc-rpc-api-user-rpc-job/pkg/broker"

	kafkaGo "github.com/segmentio/kafka-go"
)

type subscriber struct {
	k       *kafkaBroker
	topic   string
	opts    broker2.SubscribeOptions
	handler broker2.Handler
	reader  *kafkaGo.Reader
	closed  bool
	done    chan struct{}
	sync.RWMutex
}

func (s *subscriber) Options() broker2.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Topic() string {
	return s.topic
}

func (s *subscriber) Unsubscribe() error {
	var err error
	s.Lock()
	defer s.Unlock()
	s.closed = true
	return err
}
