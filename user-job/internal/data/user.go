package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "user-api-job/api/service/user-api/v1"
	"user-job/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user-api")),
	}
}

// UserById 通过rpc去查询用户信息
func (u *userRepo) UserById(c context.Context, id int64) (*biz.User, error) {
	user, err := u.data.uc.GetUserById(c, &v1.IdReq{Id: id})
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
