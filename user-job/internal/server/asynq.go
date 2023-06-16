package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	asynq2 "github.com/hibiken/asynq"
	"time"
	"user-job/internal/conf"
	"user-job/third_party/asynq"
)

func NewAsynqServer(c *conf.Bootstrap, _ log.Logger) *asynq.Server {

	srv := asynq.NewServer(
		asynq.WithAddress(c.Data.Redis.Addr),
	)

	if err := srv.HandleFunc("test_task_1", handleTask1); err != nil {
		fmt.Println(err)
	}

	if err := srv.HandleFunc("test_delay_task", handleDelayTask); err != nil {
		fmt.Println(err)
	}

	if err := srv.HandleFunc("test_periodic_task", handlePeriodicTask); err != nil {
		fmt.Println(err)
	}

	// 最多重试3次，10秒超时，20秒后过期
	err := srv.NewTask(
		"test_task_1",
		[]byte("test string"),
		asynq2.MaxRetry(10),
		asynq2.Timeout(10*time.Second),
		asynq2.Deadline(time.Now().Add(20*time.Second)),
	)

	// 延迟队列
	err = srv.NewTask("test_delay_task", []byte("delay task"), asynq2.ProcessIn(3*time.Second))

	err = srv.NewPeriodicTask("*/1 * * * ?", "test_periodic_task", []byte("periodic task"))
	if err != nil {
		fmt.Println(err)

	}

	return srv
}

func handleTask1(_ context.Context, task *asynq2.Task) error {
	fmt.Printf("Task Type: [%s], Payload: [%s]", task.Type(), string(task.Payload()))
	return nil
}

func handleDelayTask(_ context.Context, task *asynq2.Task) error {
	log.Infof("Delay Task Type: [%s], Payload: [%s]", task.Type(), string(task.Payload()))
	return nil
}

func handlePeriodicTask(_ context.Context, task *asynq2.Task) error {
	log.Infof("Periodic Task Type: [%s], Payload: [%s]", task.Type(), string(task.Payload()))
	return nil
}
