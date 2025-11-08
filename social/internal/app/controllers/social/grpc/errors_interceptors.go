package grpc

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
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
			case errors.Is(err, models.ErrSocialNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, models.ErrSocialAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, err.Error())
			case errors.Is(err, models.ErrSocialUnauthenticated):
				return nil, status.Error(codes.Unauthenticated, err.Error())
			case errors.Is(err, models.ErrSocialPermissionDenied):
				return nil, status.Error(codes.PermissionDenied, err.Error())
			case errors.Is(err, models.ErrSocialInvalidArgument):
				return nil, status.Error(codes.InvalidArgument, err.Error())
			default:
				logger.Error(ctx, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
