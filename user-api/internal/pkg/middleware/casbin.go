package middleware

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

func MiddlewareCasbin(e *casbin.Enforcer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			var uid int64

			if claims, ok := jwt.FromContext(ctx); ok {
				uid = (claims.(jwt2.MapClaims))["id"].(int64)
			} else {
				return nil, errors.Unauthorized("UNAUTHORIZED", "你没有权限！")
			}

			if tr, ok := transport.FromServerContext(ctx); ok {
				//断言成HTTP的Transport可以拿到特殊信息
				fmt.Println(tr.Operation())
				if ht, ok := tr.(*http.Transport); ok {
					ok, err := e.Enforce(uid, ht.Request().RequestURI, ht.Request().Method)
					fmt.Println(ht.Request().RequestURI, ht.Request().Method)
					if err != nil || !ok {
						// 拒绝请求，抛出异常
						return nil, errors.Unauthorized("UNAUTHORIZED", "")
					}
				}
			}

			return handler(ctx, req)
		}
	}
}
