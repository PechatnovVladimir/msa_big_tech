package grpc

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/sony/gobreaker/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Client struct {
	grpcClient *grpc.ClientConn
}

func NewGRPCClientConn(target string, cfg config.Config) (*Client, error) {
	var unaryInterceptors []grpc.UnaryClientInterceptor

	// 1. Timeout interceptor
	unaryInterceptors = append(unaryInterceptors, timeoutInterceptor(cfg.Grpc.Client.Timeout))

	// 2. Retry interceptor
	retryCodes := parseRetryCodes(cfg.Grpc.Client.Retry.RetryableCodes)

	//backoffFunc := retry.BackoffExponentialWithJitter(cfg.Grpc.Client.Retry.Backoff.Base, cfg.Grpc.Client.Retry.Backoff.Max)

	backoffFunc := buildBackoffFunc(cfg.Grpc.Client.Retry.Backoff)

	retryOpts := []retry.CallOption{
		retry.WithMax(cfg.Grpc.Client.Retry.MaxAttempts),
		retry.WithCodes(retryCodes...),
		retry.WithBackoff(backoffFunc),
	}
	unaryInterceptors = append(unaryInterceptors, retry.UnaryClientInterceptor(retryOpts...))

	// 3. Circuit Breaker (gobreaker/v2)
	cb := gobreaker.NewCircuitBreaker[string](
		gobreaker.Settings{
			Name:        fmt.Sprintf("cb-%s", target),
			Interval:    cfg.Grpc.Client.CircuitBreaker.Window,
			Timeout:     cfg.Grpc.Client.CircuitBreaker.OpenStateFor,
			MaxRequests: cfg.Grpc.Client.CircuitBreaker.HalfOpenMaxCalls,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > cfg.Grpc.Client.CircuitBreaker.FailuresForOpen
			},
		},
	)
	unaryInterceptors = append(unaryInterceptors, circuitBreakerInterceptor(cb))

	// 4. Metrics (Prometheus)
	//if cfg.Grpc.Client.Metrics {
	//	metrics := grpc_prometheus.NewClientMetrics()
	//	grpc_prometheus.Register(metrics)
	//	unaryInterceptors = append(unaryInterceptors, metrics.UnaryClientInterceptor())
	//}

	// Цепочка интерсепторов
	chainedInterceptor := chainUnaryInterceptors(unaryInterceptors...)

	// Подключение
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // В проде — TLS
		grpc.WithUnaryInterceptor(chainedInterceptor),
	)

	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: conn}, nil
}

func buildBackoffFunc(b config.BackoffConfig) retry.BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		// Базовая экспоненциальная задержка
		delay := b.Base * time.Duration(math.Pow(2, float64(attempt)))

		// Применяем jitter, если включён
		if b.Jitter {
			jitter := b.JitterFraction
			if jitter <= 0 {
				jitter = 0.5
			}
			// ±jitter от delay
			jitterDuration := time.Duration(jitter * float64(delay))
			minDelay := delay - jitterDuration
			maxDelay := delay + jitterDuration

			// Случайное значение в диапазоне
			delay = minDelay + time.Duration(rand.Float64()*float64(maxDelay-minDelay))
		}

		// Ограничиваем максимум
		if b.Max > 0 && delay > b.Max {
			delay = b.Max
		}

		return delay
	}
}

func chainUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	if len(interceptors) == 0 {
		return nil
	}
	if len(interceptors) == 1 {
		return interceptors[0]
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		currentInvoker := invoker
		// Обратный порядок: последний интерсептор — внешний
		for i := len(interceptors) - 1; i >= 0; i-- {
			interceptor := interceptors[i]
			currentInvoker = buildInterceptor(currentInvoker, interceptor)
		}
		return currentInvoker(ctx, method, req, reply, cc, opts...)
	}
}

// buildInterceptor — замыкание для цепочки
func buildInterceptor(next grpc.UnaryInvoker, interceptor grpc.UnaryClientInterceptor) grpc.UnaryInvoker {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return interceptor(ctx, method, req, reply, cc, next, opts...)
	}
}

// timeoutInterceptor sets a default deadline if none is provided in the context.
func timeoutInterceptor(defaultTimeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if _, ok := ctx.Deadline(); !ok && defaultTimeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func circuitBreakerInterceptor(cb *gobreaker.CircuitBreaker[string]) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		result, err := cb.Execute(func() (string, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return "", err
			}
			return "success", nil
		})
		if err != nil {
			return err
		}
		_ = result // игнорируем, если успех
		return nil
	}
}

// parseRetryCodes — конвертирует строки в codes.Code
func parseRetryCodes(codeStrs []string) []codes.Code {
	var codesList []codes.Code
	for _, s := range codeStrs {
		switch strings.ToUpper(strings.TrimSpace(s)) {
		case "UNAVAILABLE":
			codesList = append(codesList, codes.Unavailable)
		case "DEADLINE_EXCEEDED":
			codesList = append(codesList, codes.DeadlineExceeded)
		case "RESOURCE_EXHAUSTED":
			codesList = append(codesList, codes.ResourceExhausted)
		case "ABORTED":
			codesList = append(codesList, codes.Aborted)
		}
	}
	return codesList
}
