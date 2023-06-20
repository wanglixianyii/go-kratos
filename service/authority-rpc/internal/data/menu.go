package data

import (
	"authority-rpc/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Menu is the golang structure for table sys_menu.
type Menu struct {
	Id       int64  `orm:"id,primary"  json:"id"`       // 主键ID
	ParentId int64  `orm:"parent_id"   json:"parentId"` // 父级ID
	Title    string `orm:"title"       json:"title"`    // 菜单标题
	Icon     string `orm:"icon"        json:"icon"`     // 图标
	Path     string `orm:"path"        json:"path"`     // 菜单路径

	Permission string `orm:"permission"  json:"permission"` // 权限标识
	Type       int32  `orm:"type"        json:"type"`       // 类型：0 菜单 1 节点
	Method     string `orm:"method"      json:"method"`     // 请求方式
	Status     int32  `orm:"status"      json:"status"`     // 状态：1正常 2禁用
	Hide       int32  `orm:"hide"        json:"hide"`       // 是否可见：1是 2否
	Note       string `orm:"note"        json:"note"`       // 备注
	Sort       int32  `orm:"sort"        json:"sort"`       // 显示顺序
	CreateUser int64  `orm:"create_user" json:"createUser"` // 添加人
	CreateTime int64  `orm:"create_time" json:"createTime"` // 创建时间
	UpdateUser int64  `orm:"update_user" json:"updateUser"` // 更新人
	UpdateTime int64  `orm:"update_time" json:"updateTime"` // 更新时间

}

var _ biz.MenuRepo = (*menuRepo)(nil)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/authority-rpc")),
	}
}

func (r *menuRepo) ListMenu(ctx context.Context, title string) ([]*biz.Menu, error) {

	var menu []Menu

	result := r.data.db.WithContext(ctx).Where(&Menu{Title: title}).Find(&menu)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("MENU_NOT_FOUND", "user-rpc-rpc-api address not found")
	}

	var menuList []*biz.Menu
	for _, v := range menu {

		menuList = append(menuList, &biz.Menu{
			Id:       v.Id,
			ParentId: v.ParentId,
			Title:    v.Title,
		})
	}
	return menuList, nil
}

func (r *menuRepo) GetMenuById(ctx context.Context, id int64) (*biz.Menu, error) {
	m := Menu{}
	result := r.data.db.WithContext(ctx).First(&m, id)
	return &biz.Menu{
		Id: m.Id,
	}, result.Error
}

func (r *menuRepo) CreateMenu(ctx context.Context, b *biz.Menu) (*biz.Menu, error) {
	m := Menu{Id: b.Id, Title: b.Title}
	result := r.data.db.WithContext(ctx).Create(&m)
	return &biz.Menu{
		Id: m.Id,
	}, result.Error
}

func (r *menuRepo) UpdateMenu(ctx context.Context, b *biz.Menu) (*biz.Menu, error) {

	var menu Menu
	if result := r.data.db.First(&menu, b.Id); result.RowsAffected == 0 {
		return nil, errors.NotFound("MENU_NOT_FOUND", "菜单不存在")
	}

	if b.Title != "" {
		menu.Title = b.Title
	}

	result := r.data.db.Save(&menu)
	if result.Error != nil {
		return nil, errors.InternalServer("MENU_UPDATE_ERROR", "菜单更新失败")
	}

	var bm biz.Menu
	err := copier.Copy(&bm, &menu)
	if err != nil {
		return nil, err
	}
	return &bm, nil

}

func (r *menuRepo) DeleteMenu(ctx context.Context, id int64) error {

	result := r.data.db.WithContext(ctx).Delete(&Menu{}, id)
	return result.Error
}
