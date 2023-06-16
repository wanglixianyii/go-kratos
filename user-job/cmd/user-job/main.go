package main

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/wanglixianyii/go-kratos/user-job/pkg/bootstrap"
	"github.com/wanglixianyii/go-kratos/user-job/pkg/service"
	"github.com/wanglixianyii/go-kratos/user-job/pkg/transport/rocketmq"
	"github.com/wanglixianyii/go-kratos/user-job/third_party/asynq"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
)

var Service = bootstrap.NewServiceInfo(
	service.UserJobService,
	"1.0.0",
	"",
)

func newApp(logger log.Logger, rr registry.Registrar, rc *rocketmq.Server, as *asynq.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			rc,
			as,
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
