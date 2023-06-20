package rocketmq

import (
	aliyun "github.com/aliyunmq/mq-http-go-sdk"
	"sync"
	broker2 "user-rpc-rpc-api-user-rpc-job/pkg/broker"
)

type subscriber struct {
	r       *rocketmqBroker
	topic   string
	opts    broker2.SubscribeOptions
	handler broker2.Handler
	reader  rocketmq.PushConsumer
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

///
/// Aliyun Subscriber
///

type aliyunSubscriber struct {
	sync.RWMutex
	r       *aliyunBroker
	topic   string
	opts    broker2.SubscribeOptions
	handler broker2.Handler
	binder  broker2.Binder
	reader  aliyun.MQConsumer
	closed  bool
	done    chan struct{}
}

func (s *aliyunSubscriber) Options() broker2.SubscribeOptions {
	return s.opts
}

func (s *aliyunSubscriber) Topic() string {
	return s.topic
}

func (s *aliyunSubscriber) Unsubscribe() error {
	var err error
	s.Lock()
	defer s.Unlock()
	s.closed = true
	return err
}
