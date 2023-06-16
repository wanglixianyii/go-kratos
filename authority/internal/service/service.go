package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/wanglixianyii/go-kratos/rpc-authority/api/authority/v1"
	"github.com/wanglixianyii/go-kratos/rpc-authority/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthorityService)

type AuthorityService struct {
	v1.UnimplementedAuthorityServer

	muc *biz.MenuUseCase
	log *log.Helper
}

func NewAuthorityService(muc *biz.MenuUseCase, logger log.Logger) *AuthorityService {
	return &AuthorityService{
		muc: muc,
		log: log.NewHelper(log.With(logger, "module", "service/authority")),
	}
}
