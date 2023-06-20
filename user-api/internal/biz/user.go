package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	authJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	v1 "user-api/api/user-api/v1"

	usV1 "user-api/api/service/user-rpc/v1"
	"user-api/internal/conf"
	"user-api/internal/pkg/captcha"
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
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UserById(ctx context.Context, Id int64) (*User, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
	CreateUser(c context.Context, u *User) (*User, error)
}

type UserUseCase struct {
	repo       UserRepo
	us         usV1.UserClient
	log        *log.Helper
	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us usV1.UserClient, conf *conf.Auth) *UserUseCase {

	helper := log.NewHelper(log.With(logger, "module", "useCase/user-rpc-rpc-api"))
	return &UserUseCase{
		repo:       repo,
		us:         us,
		log:        helper,
		signingKey: conf.JwtKey,
	}
}

// GetCaptcha 验证码
func (uc *UserUseCase) GetCaptcha(ctx context.Context) (*v1.CaptchaResp, error) {

	captchaInfo, err := captcha.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.CaptchaResp{
		CaptchaId: captchaInfo.CaptchaId,
		PicPath:   captchaInfo.PicPath,
	}, nil
}

// Login 登录
func (uc *UserUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginResp, error) {

	// 表单验证
	if len(req.Mobile) <= 0 {
		return nil, errors.New("mobile invalid")
	}
	if len(req.Password) <= 0 {
		return nil, errors.New("password invalid")
	}
	if len(req.Captcha) <= 0 {
		return nil, errors.New("verification invalid")
	}

	// 验证验证码是否正确
	//if !captcha.Store.Verify(req.CaptchaId, req.Captcha, true) {
	//	return nil, errors.New("verification code error")
	//}

	user, err := uc.repo.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, errors.New("user-rpc-rpc-api not found")
	}

	// 用户存在检查密码
	passRsp, pasErr := uc.repo.CheckPassword(ctx, req.Password, user.Password)
	if pasErr != nil {
		return nil, errors.New("password invalid")
	}
	if !passRsp {
		return nil, errors.New("login failed")
	}

	// 生成token

	token, err := uc.CreateToken(user, uc.signingKey)
	if err != nil {
		return nil, errors.New("generate token failed")
	}

	return &v1.LoginResp{
		Id:        user.Id,
		Mobile:    user.Mobile,
		Username:  user.Username,
		Token:     token,
		ExpiredAt: time.Now().Unix() + 60*60*24*30,
	}, nil
}

// CreateUser 注册
func (uc *UserUseCase) CreateUser(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterResp, error) {

	user := User{
		Mobile:   req.Mobile,
		Username: req.Username,
		Password: req.Password,
	}

	createUser, err := uc.repo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	token, err := uc.CreateToken(createUser, uc.signingKey)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterResp{
		Id:        createUser.Id,
		Mobile:    createUser.Mobile,
		Username:  createUser.Username,
		Token:     token,
		ExpiredAt: time.Now().Unix() + 60*60*24*30,
	}, nil
}

func (uc *UserUseCase) CreateToken(user *User, secretKey string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Unix() + 60*60*24*30
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (uc *UserUseCase) UserDetailById(ctx context.Context) (*v1.UserInfoResp, error) {

	// 在上下文 context 中取出 claims 对象
	var uId int64
	if claims, ok := authJwt.FromContext(ctx); ok {
		c := claims.(jwt.MapClaims)
		i, ok := c["id"].(float64)
		if !ok {
			return nil, errors.New("authentication failed")
		}
		uId = int64(i)
	}

	user, err := uc.repo.UserById(ctx, uId)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoResp{
		Id:       user.Id,
		Username: user.Username,
		Mobile:   user.Mobile,
	}, nil
}
