package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
	"user-rpc/internal/biz"
	"user-rpc/internal/pkg/encryption"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         int64  `json:"id" gorm:"primary_key;column:id"`       //
	Mobile     string `json:"mobile" gorm:"column:mobile"`           // 手机号码，用户唯一标识
	Password   string `json:"password" gorm:"column:password"`       //
	Username   string `json:"username" gorm:"column:username"`       // 用户昵称
	Gender     int64  `json:"gender" gorm:"column:gender"`           // 0男1女
	AvatarId   int64  `json:"avatar_id" gorm:"column:avatar_id"`     // 头像id
	AddTime    int64  `json:"add_time" gorm:"column:add_time"`       // 添加时间
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"` // 更新时间
	DeletedAt  gorm.DeletedAt
}

func (User) TableName() string {
	return "user-rpc-rpc-api"
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 验证是否已经创建
	var user User
	result := r.data.db.WithContext(ctx).Where(&User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, errors.New(500, "USER_EXIST", "用户已存在"+u.Mobile)
	}

	encryptPassword, err := encryption.EncryptPassword(u.Password) // 密码加密
	if err != nil {
		return nil, errors.New(500, "CREAT_USER_ERROR", "密码加密失败")
	}

	user.Mobile = u.Mobile
	user.Username = u.Username
	user.Password = encryptPassword
	user.AddTime = time.Now().Unix()

	res := r.data.db.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREAT_USER_ERROR", "用户创建失败1")
	}

	var bizUser biz.User

	err = copier.Copy(&bizUser, user)
	if err != nil {
		return nil, errors.New(500, "CREAT_USER_ERROR", "赋值结构体失败")
	}

	return &bizUser, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Username: username}).WithContext(ctx).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}

func (r *userRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Mobile: mobile}).WithContext(ctx).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}

func (r *userRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Id: Id}).WithContext(ctx).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Username: username}).WithContext(ctx).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}

// ModelToResponse 转换 user-rpc-rpc-api 表中所有字段的值
func modelToResponse(user User) biz.User {
	userInfoRsp := biz.User{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		Username: user.Username,
		Gender:   user.Gender,
	}
	return userInfoRsp
}

// CheckPassword .
func (r *userRepo) CheckPassword(_ context.Context, psd, encryptedPassword string) (bool, error) {
	return encryption.EqualsPassword(psd, encryptedPassword), nil
}
