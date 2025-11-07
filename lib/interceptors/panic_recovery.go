package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime"
)

func PanicRecoveryUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			log.Printf("[PANIC] %s: %v\n%s", info.FullMethod, err, buf[:n])
			_ = status.Error(codes.Internal, "internal server error")
		}
	}()
	return handler(ctx, req)
}

func PanicRecoveryStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			log.Printf("[PANIC] %s: %v\n%s", info.FullMethod, err, buf[:n])
		}
	}()
	return handler(srv, stream)
}
