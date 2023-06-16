package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"user-job/internal/conf"
	"user-job/pkg/bootstrap"
	"user-job/pkg/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	userV1 "user-api-job/api/service/user-api/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewDiscovery, NewUserServiceClient)

// Data .
type Data struct {
	log *log.Helper
	uc  userV1.UserClient
}

// NewData .
func NewData(uc userV1.UserClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
	}
	l := log.NewHelper(log.With(logger, "module", "data/user-api-job-service"))
	return &Data{log: l, uc: uc}, cleanup, nil
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}
func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserClient {

	return userV1.NewUserClient(
		bootstrap.CreateGrpcClient(context.Background(),
			r, service.UserService,
			c.Server.Grpc.GetTimeout(),
		),
	)
}
