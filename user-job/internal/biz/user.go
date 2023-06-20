package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	usV1 "github.com/wanglixianyii/go-kratos/rpc-authority-rpc-api/api/v1"
	v1 "github.com/wanglixianyii/go-kratos/user-job/api/user-job/v1"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	UserSalt  string
	Mobile    string
	Gender    int64
	AvatarId  int64
	CreatedAt time.Time
}

type UserRepo interface {
	UserById(ctx context.Context, Id int64) (*User, error)
}

type UserUseCase struct {
	repo       UserRepo
	us         usV1.UserClient
	log        *log.Helper
	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us usV1.UserClient) *UserUseCase {

	helper := log.NewHelper(log.With(logger, "module", "rocketmq/user-rpc-rpc-api"))
	return &UserUseCase{
		repo: repo,
		us:   us,
		log:  helper,
	}
}

func (uc *UserUseCase) UserDetailById(ctx context.Context, req *v1.IdReq) (*User, error) {
	user, err := uc.repo.UserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
