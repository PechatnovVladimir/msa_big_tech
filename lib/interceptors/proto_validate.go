package interceptors

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ProtoValidateUnaryInterseptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(req.(proto.Message))
	if err != nil {
		logger.Error(ctx, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return handler(ctx, req)

}
