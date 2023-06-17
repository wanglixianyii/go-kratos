package main

import (
	"authority-api/pkg/bootstrap"
	"authority-api/pkg/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	_ "go.uber.org/automaxprocs"
)

var Service = bootstrap.NewServiceInfo(
	service.RpcAuthorityService,
	"1.0.0",
	"",
)

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	cfg, ll, reg := bootstrap.Bootstrap(Service)

	app, cleanup, err := wireApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
