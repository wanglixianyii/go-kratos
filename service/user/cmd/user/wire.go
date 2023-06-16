//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/wanglixianyii/go-kratos/user-rpc/internal/biz"
	"github.com/wanglixianyii/go-kratos/user-rpc/internal/conf"
	"github.com/wanglixianyii/go-kratos/user-rpc/internal/data"
	"github.com/wanglixianyii/go-kratos/user-rpc/internal/server"
	"github.com/wanglixianyii/go-kratos/user-rpc/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
