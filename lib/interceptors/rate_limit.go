package interceptors

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RateLimitUnaryInterceptor(cfg config.RateLimitConfig) grpc.UnaryServerInterceptor {
	globalLimiter := rate.NewLimiter(rate.Limit(cfg.ReqPerSec), cfg.ReqPerSec)

	methodLimiters := make(map[string]*rate.Limiter)
	for _, p := range cfg.Paths {
		methodLimiters[p.Path] = rate.NewLimiter(rate.Limit(p.ReqPerSec), p.ReqPerSec)
	}

	ignore := make(map[string]bool)
	for _, path := range cfg.Ignore {
		ignore[path] = true
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if ignore[info.FullMethod] {
			return handler(ctx, req)
		}

		limiter := globalLimiter
		if l, ok := methodLimiters[info.FullMethod]; ok {
			limiter = l
		}

		if err := limiter.Wait(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded: %v", err)
		}

		return handler(ctx, req)
	}
}
