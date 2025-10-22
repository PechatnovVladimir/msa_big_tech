package grpc

import (
	"context"
	"errors"
	models "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"log"

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

			switch {
			case errors.Is(err, models.ErrNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, models.ErrAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, err.Error())
			case errors.Is(err, models.ErrUnauthenticated):
				return nil, status.Error(codes.Unauthenticated, err.Error())
			case errors.Is(err, models.ErrPermissionDenied):
				return nil, status.Error(codes.PermissionDenied, err.Error())
			case errors.Is(err, models.ErrInvalidArgument):
				return nil, status.Error(codes.InvalidArgument, err.Error())
			default:
				log.Println(err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
