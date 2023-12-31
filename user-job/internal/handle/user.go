package handle

import (
	"context"
	"fmt"
	"user-job/pkg/broker"
	pb "user-rpc-rpc-api-user-rpc-job/api/user-rpc-job/v1"
)

func RegisterUserHandler(fnc UserHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *pb.IdReq:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func UserCreator() broker.Any { return &pb.IdReq{} }

type UserHandler func(_ context.Context, topic string, headers broker.Headers, msg *pb.IdReq) error
