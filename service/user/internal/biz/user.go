package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
	UserSalt string
	Gender   int64
	Mobile   string
}

type UserRepo interface {
	UserByMobile(context.Context, string) (*User, error)
	GetUserById(context.Context, int64) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
	CheckPassword(context.Context, string, string) (bool, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

func (uc *UserUseCase) UserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

func (uc *UserUseCase) UserByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.GetUserByUsername(ctx, username)
}
