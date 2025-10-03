package converter

import (
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func ListRequestsRequestProtoToDto(in *social.ListRequestsRequest) (dto.ListRequestsIN, error) {
	if in == nil {
		return dto.ListRequestsIN{}, errors.New("grpc SendFriendRequestRequest is nil")
	}

	out := dto.ListRequestsIN{
		UserID: in.UserId,
	}

	return out, nil
}

func ListRequestsResponseDtoToProto(in *dto.ListRequestsOUT) (*social.ListRequestsResponse, error) {
	if in == nil {
		return nil, errors.New("grpc is RequestsRequestDTO is nil")
	}

	out := make([]*social.FriendRequest, len(in.Requests))

	for i, request := range in.Requests {
		fr := social.FriendRequest{
			RequestId: request.RequestID,
			Status:    social.Status(request.Status),
		}
		out[i] = &fr
	}

	return &social.ListRequestsResponse{FriendRequest: out}, nil
}
