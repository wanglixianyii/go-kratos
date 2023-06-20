package data

import (
	"authority-api/internal/conf"
	"authority-api/pkg/bootstrap"
	"authority-api/pkg/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	authorityV1 "authority-api/api/service/authority-rpc/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewMenuServiceClient, NewData, NewMenuRepo, NewRedisClient)

type Data struct {
	log *log.Helper
	rdb *redis.Client
	ac  authorityV1.AuthorityClient
}

// NewData .
func NewData(redisClient *redis.Client, authorityClient authorityV1.AuthorityClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {

	}
	helperLogger := log.NewHelper(log.With(logger, "module", "authority-api-service/data"))

	return &Data{ac: authorityClient, log: helperLogger, rdb: redisClient}, cleanup, nil
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

func NewMenuServiceClient(r registry.Discovery, c *conf.Bootstrap) authorityV1.AuthorityClient {
	return authorityV1.NewAuthorityClient(bootstrap.CreateGrpcClient(context.Background(), r, service.AuthorityRpcService, c.Server.Grpc.GetTimeout()))
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/front-service"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Data.Redis.Addr,
		Password:     cfg.Data.Redis.Password,
		DB:           int(cfg.Data.Redis.Db),
		DialTimeout:  cfg.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: cfg.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.Data.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		l.Fatalf("failed opening connection to redis")
	}
	rdb.AddHook(redisotel.TracingHook{})

	return rdb
}
