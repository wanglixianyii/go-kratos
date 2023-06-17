package bootstrap

import (
	"authority-api/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

func Bootstrap(serviceInfo *ServiceInfo) (*conf.Bootstrap, log.Logger, registry.Registrar) {

	// inject command flags
	Flags := NewCommandFlags()
	Flags.Init()

	// load configs
	cfg := LoadBootstrapConfig(Flags.Conf)
	if cfg == nil {
		panic("load config failed")
	}

	// init logger
	ll := NewLoggerProvider(cfg.Logger, serviceInfo)

	// init registrar
	reg := NewRegistry(cfg.Registry)

	// init tracer
	err := NewTracerProvider(cfg.Tracer, serviceInfo)
	if err != nil {
		panic(err)
	}

	return cfg, ll, reg

}
