package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func fromListRequestsRequestToDto(in *social.ListRequestsRequest) *dto.ListRequestsIN {
	return &dto.ListRequestsIN{
		UserID: in.UserId,
	}
}

func fromDtoToListRequestsResponse(in *dto.ListRequestsOUT) *social.ListRequestsResponse {
	requests := make([]*social.FriendRequest, len(in.Requests))
	for i, _ := range in.Requests {
		requests[i] = &social.FriendRequest{
			RequestId: in.Requests[i].RequestID,
			Status:    social.Status(in.Requests[i].Status),
		}
	}

	return &social.ListRequestsResponse{
		FriendRequest: requests,
	}
}
