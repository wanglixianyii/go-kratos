package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/rocketmq"
	"user-api/internal/conf"
	"user-api/internal/service"
)

func NewRocketmqServer(c *conf.Server, svc *service.LoggerJobService, _ log.Logger) *rocketmq.Server {

	srv := rocketmq.NewServer(
		rocketmq.WithNameServer(c.Rocketmq.Addrs),
		rocketmq.WithCodec("json"),
	)

	ctx := context.Background()

	_ = srv.RegisterSubscriber(
		ctx,
		"SimpleTopic",
		"simple_test",
		registerSensorHandler(svc.InsertSensor),
		sensorCreator,
	)

	return srv
}
