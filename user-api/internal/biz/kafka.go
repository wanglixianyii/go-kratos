package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	svcV1 "github.com/wanglixianyii/go-kratos/user-api/api/service/job/v1"
)

type KafkaRepo interface {
	CreateSensor(ctx context.Context, req *svcV1.Sensor) error
}

type KafkaUseCase struct {
	repo KafkaRepo
	log  *log.Helper
}

func NewKafkaUseCase(repo KafkaRepo, logger log.Logger) *KafkaUseCase {

	helper := log.NewHelper(log.With(logger, "module", "useCase/job"))
	return &KafkaUseCase{
		repo: repo,
		log:  helper,
	}
}

func (uc *KafkaUseCase) Create(ctx context.Context, req *svcV1.Sensor) error {
	return uc.repo.CreateSensor(ctx, req)
}
