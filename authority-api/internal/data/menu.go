package data

import (
	v1 "authority-api/api/service/authority-rpc/v1"
	"authority-api/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.MenuRepo = (*menuRepo)(nil)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "authority-api/data")),
	}
}

func (r *menuRepo) ListMenu(ctx context.Context, title string) (*v1.MenuListResp, error) {

	return r.data.ac.MenuList(ctx, &v1.MenuListReq{Title: title})

}

func (r *menuRepo) GetMenuById(ctx context.Context, id int64) (*v1.DetailMenuResp, error) {
	return r.data.ac.DetailMenu(ctx, &v1.DetailMenuReq{Id: id})
}

func (r *menuRepo) CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (*v1.CreateMenuResp, error) {
	return r.data.ac.CreateMenu(ctx, req)
}

func (r *menuRepo) UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (*v1.UpdateMenuResp, error) {
	return r.data.ac.UpdateMenu(ctx, req)
}

func (r *menuRepo) DeleteMenu(ctx context.Context, id int64) (*v1.DeleteMenuResp, error) {

	return r.data.ac.DeleteMenu(ctx, &v1.DeleteMenuReq{Id: id})
}
