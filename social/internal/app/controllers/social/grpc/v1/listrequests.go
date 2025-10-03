package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) ListRequests(ctx context.Context, request *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	log.Println("SocialService ListRequests called")

	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inUC, err := converter.ListRequestsRequestProtoToDto(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	outUC, err := s.SocialUseCase.ListRequests(ctx, inUC)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := converter.ListRequestsResponseDtoToProto(&outUC)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return out, nil
}
