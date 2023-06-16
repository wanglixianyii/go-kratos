package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wanglixianyii/go-kratos/user-job/internal/conf"
	"github.com/wanglixianyii/go-kratos/user-job/internal/handle"
	"github.com/wanglixianyii/go-kratos/user-job/internal/service"
	"github.com/wanglixianyii/go-kratos/user-job/pkg/transport/rocketmq"
)

func NewRocketmqServer(c *conf.Bootstrap, svc *service.UserService, _ log.Logger) *rocketmq.Server {

	srv := rocketmq.NewServer(
		rocketmq.WithNameServer(c.Server.Rocketmq.Addrs),
		rocketmq.WithCodec("json"),
	)

	ctx := context.Background()

	_ = srv.RegisterSubscriber(
		ctx,
		"SimpleTopic",
		"simple_test",
		handle.RegisterUserHandler(svc.UserInfo),
		handle.UserCreator,
	)

	return srv
}
