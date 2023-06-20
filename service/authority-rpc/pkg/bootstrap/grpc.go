package bootstrap

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	kratosGrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

const defaultTimeout = 5 * time.Second

// CreateGrpcClient 创建GRPC客户端
func CreateGrpcClient(ctx context.Context, r registry.Discovery, serviceName string, timeoutDuration *durationpb.Duration) grpc.ClientConnInterface {
	timeout := defaultTimeout
	if timeoutDuration != nil {
		timeout = timeoutDuration.AsDuration()
	}

	endpoint := "discovery:///" + serviceName

	conn, err := kratosGrpc.DialInsecure(
		ctx,
		kratosGrpc.WithEndpoint(endpoint),
		kratosGrpc.WithDiscovery(r),
		kratosGrpc.WithTimeout(timeout),
		kratosGrpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}

	return conn
}
