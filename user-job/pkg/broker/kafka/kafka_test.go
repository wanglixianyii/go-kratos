package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
	"user-job/pkg/tracing"
	broker2 "user-rpc-rpc-api-user-rpc-job/pkg/broker"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/stretchr/testify/assert"

	api "user-rpc-rpc-api-user-rpc-job/internal/pkg/_example/api/manual"
)

const (
	testBrokers = "localhost:9092"
	testTopic   = "logger.sensor.ts"
	testGroupId = "fx-group"
)

func handleHygrothermograph(_ context.Context, topic string, headers broker2.Headers, msg *api.Hygrothermograph) error {
	log.Infof("Topic %s, Headers: %+v, Payload: %+v\n", topic, headers, msg)
	return nil
}

func Test_Publish_WithRawData(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
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
		buf, _ := json.Marshal(&msg)
		err := b.Publish(testTopic, buf)
		assert.Nil(t, err)
		elapsedTime := time.Since(startTime) / time.Millisecond
		fmt.Printf("Publish %d, elapsed time: %dms, Humidity: %.2f Temperature: %.2f\n",
			i, elapsedTime, msg.Humidity, msg.Temperature)
	}

	fmt.Printf("total send %d messages\n", count)

	<-interrupt
}

func Test_Subscribe_WithRawData(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	_, err := b.Subscribe(testTopic,
		api.RegisterHygrothermographRawHandler(handleHygrothermograph),
		nil,
		broker2.WithQueueName(testGroupId),
	)
	assert.Nil(t, err)
	assert.Nil(t, err)

	<-interrupt
}

func Test_Publish_WithJsonCodec(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
		broker2.WithCodec("json"),
		//WithAsync(false),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	var headers map[string]interface{}
	headers = make(map[string]interface{})
	headers["version"] = "1.0.0"

	var msg api.Hygrothermograph
	const count = 10
	for i := 0; i < count; i++ {
		startTime := time.Now()
		headers["trace_id"] = i
		msg.Humidity = float64(rand.Intn(100))
		msg.Temperature = float64(rand.Intn(100))
		err := b.Publish(testTopic, msg, WithHeaders(headers))
		assert.Nil(t, err)
		elapsedTime := time.Since(startTime) / time.Millisecond
		log.Infof("Publish %d, elapsed time: %dms, Humidity: %.2f Temperature: %.2f\n",
			i, elapsedTime, msg.Humidity, msg.Temperature)
	}

	log.Infof("total send %d messages\n", count)

	<-interrupt
}

func Test_Subscribe_WithJsonCodec(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
		broker2.WithCodec("json"),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	_, err := b.Subscribe(testTopic,
		api.RegisterHygrothermographJsonHandler(handleHygrothermograph),
		api.HygrothermographCreator,
		broker2.WithQueueName(testGroupId),
	)
	assert.Nil(t, err)

	<-interrupt
}

func createTracerProvider(exporterName, serviceName string) broker2.Option {
	switch exporterName {
	case "jaeger":
		return broker2.WithTracerProvider(tracing.NewTracerProvider(exporterName,
			"http://localhost:14268/api/traces",
			serviceName,
			"",
			"1.0.0",
			1.0,
		),
			"kafka-tracer",
		)
	case "zipkin":
		return broker2.WithTracerProvider(tracing.NewTracerProvider(exporterName,
			"http://localhost:9411/api/v2/spans",
			serviceName,
			"test",
			"1.0.0",
			1.0,
		),
			"kafka-tracer",
		)
	}

	return nil
}

func Test_Publish_WithTracer(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
		broker2.WithCodec("json"),
		createTracerProvider("jaeger", "tracer_tester"),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	var msg api.Hygrothermograph
	const count = 1
	for i := 0; i < count; i++ {
		startTime := time.Now()
		msg.Humidity = float64(rand.Intn(100))
		msg.Temperature = float64(rand.Intn(100))
		err := b.Publish(testTopic, msg)
		assert.Nil(t, err)
		elapsedTime := time.Since(startTime) / time.Millisecond
		log.Infof("Publish %d, elapsed time: %dms, Humidity: %.2f Temperature: %.2f\n",
			i, elapsedTime, msg.Humidity, msg.Temperature)
	}

	log.Infof("total send %d messages\n", count)

	<-interrupt
}

func Test_Subscribe_WithTracer(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := NewBroker(
		broker2.WithAddress(testBrokers),
		broker2.WithCodec("json"),
		createTracerProvider("jaeger", "subscribe_tracer_tester"),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}

	_, err := b.Subscribe(testTopic,
		api.RegisterHygrothermographJsonHandler(handleHygrothermograph),
		api.HygrothermographCreator,
		broker2.WithQueueName(testGroupId),
	)
	assert.Nil(t, err)

	<-interrupt
}
