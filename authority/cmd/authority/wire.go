//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/biz"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/conf"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/data"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/server"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
