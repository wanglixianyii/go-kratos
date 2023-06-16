package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"user-job/internal/biz"
	"user-job/pkg/broker"

	pb "user-api-job/api/job/v1"
	v1 "user-api-job/api/service/user-api/v1"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user-api")),
	}
}

func (s *UserService) UserInfo(ctx context.Context, topic string, _ broker.Headers, msg *pb.IdReq) error {
	fmt.Println("UserInfo() Topic: ", topic)

	user, err := s.uc.UserDetailById(ctx, msg)
	if err != nil {
		s.log.Debug("get UserInfo", err.Error())

		return err
	}
	fmt.Println(user)
	return nil
}
