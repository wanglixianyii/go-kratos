package service

import (
	v1 "authority-api/api/service/authority-rpc/v1"
	"authority-api/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuService struct {
	v1.UnimplementedAuthorityServer

	muc *biz.MenuUseCase
	log *log.Helper
}

func NewMenuService(muc *biz.MenuUseCase, logger log.Logger) *MenuService {
	return &MenuService{
		muc: muc,
		log: log.NewHelper(log.With(logger, "module", "authority-api/service")),
	}
}

func (s *MenuService) MenuList(ctx context.Context, req *v1.MenuListReq) (*v1.MenuListResp, error) {

	return s.muc.List(ctx, req)
}

func (s *MenuService) CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (*v1.CreateMenuResp, error) {

	return s.muc.Create(ctx, req)
}

func (s *MenuService) UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (*v1.UpdateMenuResp, error) {

	return s.muc.Update(ctx, req)
}

func (s *MenuService) DeleteMenu(ctx context.Context, req *v1.DeleteMenuReq) (*v1.DeleteMenuResp, error) {

	return s.muc.Delete(ctx, req)
}

func (s *MenuService) DetailMenu(ctx context.Context, req *v1.DetailMenuReq) (*v1.DetailMenuResp, error) {

	return s.muc.GetMenuById(ctx, req.Id)
}
