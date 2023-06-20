package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-api/api/user-api/v1"
	"user-api/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewAdminService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user-rpc-rpc-api")),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user-rpc-rpc-api info by mobile")
	span.SpanContext()
	defer span.End()

	return s.uc.CreateUser(ctx, req)
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	s.log.Info("123")
	return s.uc.Login(ctx, req)
}
func (s *UserService) Captcha(ctx context.Context, _ *emptypb.Empty) (*pb.CaptchaResp, error) {
	return s.uc.GetCaptcha(ctx)
}
func (s *UserService) UserInfo(ctx context.Context, _ *emptypb.Empty) (*pb.UserInfoResp, error) {
	return s.uc.UserDetailById(ctx)
}
