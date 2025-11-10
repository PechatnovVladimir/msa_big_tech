package interceptors

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/tracing"
	//"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func OtelUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Создаем корневой span
	ctx, span := tracing.Start(ctx, "grpc "+info.FullMethod, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	// Вызываем следующий обработчик (или сам handler)
	resp, err := handler(ctx, req)

	// Записываем полезные атрибуты
	span.SetAttributes(
		semconv.HTTPRoute(info.FullMethod),
	)

	// Помечаем span как ошибочный для 4xx и 5xx статусов
	if err != nil {
		span.SetStatus(codes.Error, "")
		span.AddEvent("error", trace.WithAttributes(
			attribute.String("error.message", err.Error()),
		))
	}

	return resp, err
}
