package server

import (
	svcV1 "user-api/api/service/job/v1"
	"admin/internal/conf"
	"admin/internal/service"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/kafka"
)

func NewKafKaServer(c *conf.Server, svc *service.LoggerJobService, _ log.Logger) *kafka.Server {

	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv

}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.LoggerJobService) {

	_ = srv.RegisterSubscriber(ctx,
		"logger.sensor.instance",
		"sensor",
		false,
		registerSensorHandler(svc.InsertSensor),
		sensorCreator,
	)
}

type SensorHandler func(_ context.Context, topic string, headers broker.Headers, msg *svcV1.Sensor) error

func registerSensorHandler(fnc SensorHandler) broker.Handler {

	return func(ctx context.Context, event broker.Event) error {
		fmt.Println("2")
		fmt.Println(event.Message().Headers)
		switch t := event.Message().Body.(type) {
		case *svcV1.Sensor:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}

}
func sensorCreator() broker.Any { return &svcV1.Sensor{} }
