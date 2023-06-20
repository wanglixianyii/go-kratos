package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	userService "user-api/api/service/user-rpc/v1"
	"user-api/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user-rpc-rpc-api")),
	}
}

// CreateUser 通过rpc去添加用户信息
func (u *userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := u.data.uc.CreateUser(c, &userService.CreateUserReq{
		Username: user.Username,
		Password: user.Password,
		Mobile:   user.Mobile,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       createUser.Id,
		Mobile:   createUser.Mobile,
		Username: createUser.Username,
	}, nil
}

// UserByMobile 通过rpc去查询用户信息
func (u *userRepo) UserByMobile(c context.Context, mobile string) (*biz.User, error) {
	byMobile, err := u.data.uc.GetUserByMobile(c, &userService.MobileReq{Mobile: mobile})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &biz.User{

		Id:       byMobile.Id,
		Password: byMobile.Password,
		Username: byMobile.Username,
		Mobile:   byMobile.Mobile,
	}, nil
}

// UserById 通过rpc去查询用户信息
func (u *userRepo) UserById(c context.Context, id int64) (*biz.User, error) {
	user, err := u.data.uc.GetUserById(c, &userService.IdReq{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Username: user.Username,
		Gender:   user.Gender,
		AvatarId: user.AvatarId,
	}, nil
}

// UserById 通过rpc去验证密码
func (u *userRepo) CheckPassword(c context.Context, password, encryptedPassword string) (bool, error) {
	if byMobile, err := u.data.uc.CheckPassword(c, &userService.CheckPasswordReq{Password: password, EncryptedPassword: encryptedPassword}); err != nil {
		return false, err
	} else {
		return byMobile.Success, nil
	}
}
