//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/wanglixianyii/go-kratos/user-api/internal/biz"
	"github.com/wanglixianyii/go-kratos/user-api/internal/conf"
	"github.com/wanglixianyii/go-kratos/user-api/internal/data"
	"github.com/wanglixianyii/go-kratos/user-api/internal/server"
	"github.com/wanglixianyii/go-kratos/user-api/internal/service"
)

func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Service, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
