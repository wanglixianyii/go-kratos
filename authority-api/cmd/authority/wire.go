//go:build wireinject
// +build wireinject

package main

import (
	"authority-api/internal/biz"
	"authority-api/internal/conf"
	"authority-api/internal/data"
	"authority-api/internal/server"
	"authority-api/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
