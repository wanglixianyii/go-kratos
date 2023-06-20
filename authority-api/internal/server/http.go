package server

import (
	v1 "authority-api/api/authority-api/v1"
	"authority-api/internal/conf"
	"authority-api/internal/service"
	"context"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	midRateLimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, s *service.MenuService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders(c.Server.Http.Cors.Headers),
			handlers.AllowedMethods(c.Server.Http.Cors.Methods),
			handlers.AllowedOrigins(c.Server.Http.Cors.Origins),
		)),
	}
	var ms = []middleware.Middleware{
		recovery.Recovery(),
		validate.Validator(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(c.Server.Http.Middleware.Auth.JwtKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		).Match(NewWhiteListMatcher()).Build(),
	}
	if c.Server.Http.Middleware.Limiter != nil {
		var limiter ratelimit.Limiter
		switch c.Server.Http.Middleware.Limiter.GetName() {
		case "bbr":
			limiter = bbr.NewLimiter()
		}
		ms = append(ms, midRateLimit.Server(midRateLimit.WithLimiter(limiter)))
	}

	opts = append(opts, http.Middleware(ms...))

	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterAuthorityHTTPServer(srv, s)
	return srv
}

// NewWhiteListMatcher 白名单不需要token验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthorityMenuList] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
