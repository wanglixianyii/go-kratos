package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	v1 "user-rpc/api/user-rpc/v1"
	"user-rpc/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.UserInfoResp, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := v1.UserInfoResp{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		Username: user.Username,
		Gender:   user.Gender,
	}

	return &userInfoRsp, nil
}
func (s *UserService) CheckPassword(ctx context.Context, req *v1.CheckPasswordReq) (*v1.CheckPasswordResp, error) {

	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "check user-rpc-rpc-api  password")
	defer span.End()
	check, err := s.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.CheckPasswordResp{Success: check}, nil

}
func (s *UserService) GetUserByMobile(ctx context.Context, req *v1.MobileReq) (*v1.UserInfoResp, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user-rpc-rpc-api list")
	defer span.End()
	user, err := s.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	var resp v1.UserInfoResp
	err = copier.Copy(&resp, user)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.UsernameReq) (*v1.UserInfoResp, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user-rpc-rpc-api info by Id")
	defer span.End()
	user, err := s.uc.UserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	var resp v1.UserInfoResp
	err = copier.Copy(&resp, user)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *v1.IdReq) (*v1.UserInfoResp, error) {
	//tr := otel.Tracer("service")
	//ctx, span := tr.Start(ctx, "get user-rpc-rpc-api info by Id")
	//defer span.End()
	user, err := s.uc.UserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := UserResponse(user)

	s.log.Info("12q")
	return &rsp, nil
}

func UserResponse(user *biz.User) v1.UserInfoResp {
	userInfoRsp := v1.UserInfoResp{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		Username: user.Username,
		Gender:   user.Gender,
	}

	return userInfoRsp
}
