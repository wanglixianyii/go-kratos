package rocketmq

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
	broker2 "user-rpc-rpc-api-user-rpc-job/pkg/broker"
	rocketmq2 "user-rpc-rpc-api-user-rpc-job/pkg/broker/rocketmq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	api "user-rpc-rpc-api-user-rpc-job/internal/pkg/_example/api/manual"
)

const (
	// Name Server address
	testBroker = "127.0.0.1:9876"

	testTopic     = "test_topic"
	testGroupName = "CID_ONSAPI_OWNER"
)

func handleHygrothermograph(_ context.Context, topic string, headers broker2.Headers, msg *api.Hygrothermograph) error {
	log.Infof("Topic %s, Headers: %+v, Payload: %+v\n", topic, headers, msg)
	return nil
}

func TestServer(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()

	srv := NewServer(
		WithNameServer([]string{testBroker}),
		//WithNameServerDomain("http://nsaddr.rmq.cloud.tencent.com"),
		WithCodec("json"),
	)

	err := srv.RegisterSubscriber(ctx, testTopic, testGroupName,
		api.RegisterHygrothermographJsonHandler(handleHygrothermograph),
		api.HygrothermographCreator,
	)
	assert.Nil(t, err)

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()

	<-interrupt
}

func TestClient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := rocketmq2.NewBroker(
		broker2.WithCodec("json"),
		rocketmq2.WithEnableTrace(),
		rocketmq2.WithNameServer([]string{testBroker}),
		//rocketmq.WithNameServerDomain(testBroker),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	var msg api.Hygrothermograph
	const count = 10
	for i := 0; i < count; i++ {
		startTime := time.Now()
		msg.Humidity = float64(rand.Intn(100))
		msg.Temperature = float64(rand.Intn(100))
		err := b.Publish(testTopic, msg)
		assert.Nil(t, err)
		elapsedTime := time.Since(startTime) / time.Millisecond
		fmt.Printf("Publish %d, elapsed time: %dms, Humidity: %.2f Temperature: %.2f\n",
			i, elapsedTime, msg.Humidity, msg.Temperature)
	}

	fmt.Printf("total send %d messages\n", count)

	<-interrupt
}

func TestAliyunServer(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	endpoint := ""
	accessKey := ""
	secretKey := ""
	instanceId := ""
	topicName := ""
	groupName := "GID_DEFAULT"

	srv := NewServer(
		WithAliyunHttpSupport(),
		WithCodec("json"),
		WithEnableTrace(),
		WithNameServerDomain(endpoint),
		WithCredentials(accessKey, secretKey, ""),
		WithInstanceName(instanceId),
		WithGroupName(groupName),
	)

	err := srv.RegisterSubscriber(ctx, topicName, groupName,
		api.RegisterHygrothermographJsonHandler(handleHygrothermograph),
		api.HygrothermographCreator,
	)
	assert.Nil(t, err)

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()

	<-interrupt
}

func TestAliyunClient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	endpoint := ""
	accessKey := ""
	secretKey := ""
	instanceId := ""
	topicName := ""

	b := rocketmq2.NewBroker(
		broker2.WithOptionContext(ctx),
		broker2.WithCodec("json"),
		rocketmq2.WithAliyunHttpSupport(),
		rocketmq2.WithEnableTrace(),
		rocketmq2.WithNameServerDomain(endpoint),
		rocketmq2.WithAccessKey(accessKey),
		rocketmq2.WithSecretKey(secretKey),
		rocketmq2.WithInstanceName(instanceId),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	var msg api.Hygrothermograph
	const count = 10
	for i := 0; i < count; i++ {
		startTime := time.Now()
		msg.Humidity = float64(rand.Intn(100))
		msg.Temperature = float64(rand.Intn(100))
		err := b.Publish(topicName, msg)
		assert.Nil(t, err)
		elapsedTime := time.Since(startTime) / time.Millisecond
		fmt.Printf("Publish %d, elapsed time: %dms, Humidity: %.2f Temperature: %.2f\n",
			i, elapsedTime, msg.Humidity, msg.Temperature)
	}

	fmt.Printf("total send %d messages\n", count)

	<-interrupt
}
