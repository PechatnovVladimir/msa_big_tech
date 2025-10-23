package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorsUnaryServerInterceptor - convert any error to rpc error
func ErrorsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}

			return nil, status.Error(codes.Internal, err.Error())
		}

		return
	}
}
