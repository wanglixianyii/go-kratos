package data

import (
	svcV1 "user-api/api/service/job/v1"
	"admin/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.KafkaRepo = (*kafkaRepo)(nil)

type kafkaRepo struct {
	data *Data
	log  *log.Helper
}

func NewKafkaRepo(data *Data, logger log.Logger) biz.KafkaRepo {
	return &kafkaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "job/repo/logger-service")),
	}
}

func (r kafkaRepo) CreateSensor(_ context.Context, req *svcV1.Sensor) error {
	fmt.Println("req", req)
	return nil
}
