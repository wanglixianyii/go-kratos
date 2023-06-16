package service

import (
	v1 "user-api/api/job/v1"
	svcV1 "user-api/api/service/job/v1"
	"admin/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

type LoggerJobService struct {
	v1.UnimplementedLoggerJobServer

	log *log.Helper
	uc  *biz.KafkaUseCase
}

func NewLoggerJobService(uc *biz.KafkaUseCase, logger log.Logger) *LoggerJobService {
	return &LoggerJobService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/logger-job"))}
}

func (s *LoggerJobService) InsertSensor(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.Sensor) error {
	fmt.Println("InsertSensor() Topic: ", topic)

	if err := s.uc.Create(context.Background(), msg); err != nil {
		s.log.Debug("InsertSensor Insert", err.Error())
		return err
	}

	return nil
}
