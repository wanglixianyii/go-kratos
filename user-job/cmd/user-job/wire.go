//go:build wireinject
// +build wireinject

//
//The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"user-job/internal/biz"
	"user-job/internal/conf"
	"user-job/internal/data"
	"user-job/internal/server"
	"user-job/internal/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
