package biz

import (
	v1 "authority-api/api/service/authority-rpc/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Menu struct {
	Id         int64  `json:"id"`         // 主键ID
	ParentId   int64  `json:"parentId"`   // 父级ID
	Title      string `json:"title"`      // 菜单标题
	Icon       string `json:"icon"`       // 图标
	Path       string `json:"path"`       // 菜单路径
	Permission string `json:"permission"` // 权限标识
	Type       int32  `json:"type"`       // 类型：0 菜单 1 节点
	Method     string `json:"method"`     // 请求方式
	Status     int32  `json:"status"`     // 状态：1正常 2禁用
	Hide       int32  `json:"hide"`       // 是否可见：1是 2否
	Note       string `json:"note"`       // 备注
	Sort       int32  `json:"sort"`       // 显示顺序
	CreateUser int64  `json:"createUser"` // 添加人
	CreateTime int64  `json:"createTime"` // 创建时间
	UpdateUser int64  `json:"updateUser"` // 更新人
	UpdateTime int64  `json:"updateTime"` // 更新时间
}

type MenuRepo interface {
	GetMenuById(ctx context.Context, id int64) (*v1.DetailMenuResp, error)
	CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (*v1.CreateMenuResp, error)
	UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (*v1.UpdateMenuResp, error)
	DeleteMenu(ctx context.Context, id int64) (*v1.DeleteMenuResp, error)
	ListMenu(ctx context.Context, title string) (*v1.MenuListResp, error)
}

type MenuUseCase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUseCase(repo MenuRepo, logger log.Logger) *MenuUseCase {
	return &MenuUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "useCase/authority-rpc-api"))}
}

func (uc *MenuUseCase) List(ctx context.Context, req *v1.MenuListReq) (*v1.MenuListResp, error) {
	uc.log.Info("test")

	return uc.repo.ListMenu(ctx, req.Title)
}

func (uc *MenuUseCase) Create(ctx context.Context, req *v1.CreateMenuReq) (*v1.CreateMenuResp, error) {
	return uc.repo.CreateMenu(ctx, req)
}

func (uc *MenuUseCase) GetMenuById(ctx context.Context, id int64) (*v1.DetailMenuResp, error) {
	return uc.repo.GetMenuById(ctx, id)
}

func (uc *MenuUseCase) Update(ctx context.Context, req *v1.UpdateMenuReq) (*v1.UpdateMenuResp, error) {
	return uc.repo.UpdateMenu(ctx, req)
}

func (uc *MenuUseCase) Delete(ctx context.Context, req *v1.DeleteMenuReq) (*v1.DeleteMenuResp, error) {
	return uc.repo.DeleteMenu(ctx, req.Id)
}
