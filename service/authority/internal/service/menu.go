package service

import (
	v1 "authority-rpc/api/authority/v1"
	"authority-rpc/internal/biz"
	"context"
)

func (s *AuthorityService) MenuList(ctx context.Context, req *v1.MenuListReq) (*v1.MenuListResp, error) {

	resp := &v1.MenuListResp{List: make([]*v1.MenuInfo, 0)}

	c, err := s.muc.List(ctx, req.Title)
	if err != nil {
		return resp, err
	}
	for _, x := range c {
		resp.List = append(resp.List,
			&v1.MenuInfo{
				Id:    x.Id,
				Title: x.Title,
			})
	}
	return resp, nil
}

func (s *AuthorityService) CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (*v1.CreateMenuResp, error) {

	b := &biz.Menu{
		Title: req.Title,
		Icon:  req.Icon,
		Path:  req.Path,
		Type:  req.Type,
		Hide:  req.Hide,
		Note:  req.Note,
		Sort:  req.Sort,
	}

	x, err := s.muc.Create(ctx, b)
	if err != nil {
		return nil, err
	}

	return &v1.CreateMenuResp{Id: x.Id}, err
}

func (s *AuthorityService) UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (*v1.UpdateMenuResp, error) {

	b := &biz.Menu{
		Title: req.Title,
		Icon:  req.Icon,
		Path:  req.Path,
		Type:  req.Type,
		Hide:  req.Hide,
		Note:  req.Note,
		Sort:  req.Sort,
	}

	x, err := s.muc.Update(ctx, b)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMenuResp{Id: x.Id}, err
}

func (s *AuthorityService) DeleteMenu(ctx context.Context, req *v1.DeleteMenuReq) (*v1.DeleteMenuResp, error) {

	err := s.muc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteMenuResp{}, err
}

func (s *AuthorityService) DetailMenu(ctx context.Context, req *v1.DetailMenuReq) (*v1.DetailMenuResp, error) {

	x, err := s.muc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DetailMenuResp{Info: &v1.MenuInfo{Id: x.Id}}, err
}
