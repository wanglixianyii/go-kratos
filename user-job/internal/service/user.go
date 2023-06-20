package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"user-job/internal/biz"
	"user-job/pkg/broker"

	v1 "user-rpc-rpc-api-user-rpc-job/api/service/user-rpc-rpc-api/v1"
	pb "user-rpc-rpc-api-user-rpc-job/api/user-rpc-job/v1"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user-rpc-rpc-api")),
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
