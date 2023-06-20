package data

import (
	"admin/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	svcV1 "user-api/api/service/job/v1"
)

var _ biz.KafkaRepo = (*kafkaRepo)(nil)

type kafkaRepo struct {
	data *Data
	log  *log.Helper
}

func NewKafkaRepo(data *Data, logger log.Logger) biz.KafkaRepo {
	return &kafkaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "user-rpc-job/repo/logger-service")),
	}
}

func (r kafkaRepo) CreateSensor(_ context.Context, req *svcV1.Sensor) error {
	fmt.Println("req", req)
	return nil
}
