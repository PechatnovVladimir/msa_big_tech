package interceptors

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"google.golang.org/grpc"
	"time"
)

func TimeoutUnaryInterceptor(cfg config.TimeoutConfig) grpc.UnaryServerInterceptor {
	methodTimeouts := make(map[string]time.Duration)
	for _, p := range cfg.Paths {
		methodTimeouts[p.Path] = time.Duration(p.TimeoutMs) * time.Millisecond
	}
	defaultTimeout := time.Duration(cfg.TimeoutMs) * time.Millisecond

	ignore := make(map[string]bool)
	for _, path := range cfg.Ignore {
		ignore[path] = true
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if ignore[info.FullMethod] {
			return handler(ctx, req)
		}

		timeout := defaultTimeout
		if t, ok := methodTimeouts[info.FullMethod]; ok && t > 0 {
			timeout = t
		}

		if timeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}

		return handler(ctx, req)
	}
}
