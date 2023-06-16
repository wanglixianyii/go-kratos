package service

import (
	pb "user-api/api/user-api/v1"
	"admin/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewAdminService(uc *biz.UserUseCase, logger log.Logger) *AdminService {
	return &AdminService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user-api")),
	}
}

func (s *AdminService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user-api info by mobile")
	span.SpanContext()
	defer span.End()

	return s.uc.CreateUser(ctx, req)
}
func (s *AdminService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	s.log.Info("123")
	return s.uc.Login(ctx, req)
}
func (s *AdminService) Captcha(ctx context.Context, _ *emptypb.Empty) (*pb.CaptchaResp, error) {
	return s.uc.GetCaptcha(ctx)
}
func (s *AdminService) UserInfo(ctx context.Context, _ *emptypb.Empty) (*pb.UserInfoResp, error) {
	return s.uc.UserDetailById(ctx)
}
