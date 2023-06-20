//go:build wireinject
// +build wireinject

//The build tag makes sure the stub is not built in the final build.

package main

import (
	"authority-rpc/internal/biz"
	"authority-rpc/internal/conf"
	"authority-rpc/internal/data"
	"authority-rpc/internal/server"
	"authority-rpc/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
