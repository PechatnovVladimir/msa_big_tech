package interceptors

import (
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServerInterceptors(cfg config.GrpcServerConfig) []grpc.ServerOption {
	var unaryInterceptors []grpc.UnaryServerInterceptor
	var streamInterceptors []grpc.StreamServerInterceptor

	//валидация
	unaryInterceptors = append(unaryInterceptors, ProtoValidateUnaryInterseptor)

	//tracing
	unaryInterceptors = append(unaryInterceptors, OtelUnaryInterceptor)

	//log errors
	unaryInterceptors = append(unaryInterceptors, LogErrorUnaryInterceptor())

	//panic recovery
	unaryInterceptors = append(unaryInterceptors, PanicRecoveryUnaryInterceptor)
	streamInterceptors = append(streamInterceptors, PanicRecoveryStreamInterceptor)

	//Timeout
	if cfg.Timeout.Enabled {
		unaryInterceptors = append(unaryInterceptors, TimeoutUnaryInterceptor(cfg.Timeout))
	}

	//Rate Limit
	if cfg.RateLimit.Enabled {
		unaryInterceptors = append(unaryInterceptors, RateLimitUnaryInterceptor(cfg.RateLimit))
	}

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
		grpc.ChainStreamInterceptor(streamInterceptors...),
		grpc.Creds(insecure.NewCredentials()),
	}
}
