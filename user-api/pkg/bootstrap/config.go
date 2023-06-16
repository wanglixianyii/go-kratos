package bootstrap

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"user-job/internal/conf"
)

// NewConfigProvider 创建一个配置
func NewConfigProvider(configPath string) config.Config {
	return config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)

}

// LoadBootstrapConfig 加载程序引导配置
func LoadBootstrapConfig(configPath string) *conf.Bootstrap {
	cfg := NewConfigProvider(configPath)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server == nil {
		bc.Server = &conf.Server{}
		_ = cfg.Scan(&bc.Server)
	}

	if bc.Data == nil {
		bc.Data = &conf.Data{}
		_ = cfg.Scan(&bc.Data)
	}

	if bc.Tracer == nil {
		bc.Tracer = &conf.Tracer{}
		_ = cfg.Scan(&bc.Tracer)
	}

	if bc.Logger == nil {
		bc.Logger = &conf.Logger{}
		_ = cfg.Scan(&bc.Logger)
	}

	if bc.Registry == nil {
		bc.Registry = &conf.Registry{}
		_ = cfg.Scan(&bc.Registry)
	}
	return &bc
}
