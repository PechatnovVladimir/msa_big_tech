package converter

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"golang.org/x/net/context"
)

func ListRequestsFromRepoToModel(ctx context.Context, in dtoRepo.ListRequestsOUT) (models.FriendRequests, error) {

	out := make(models.FriendRequests, len(in.FriendRequests))

	for i, friendRequest := range in.FriendRequests {
		out[i].RequestID = friendRequest.RequestID
		out[i].FromUserID = friendRequest.FromUserID
		out[i].ToUserID = friendRequest.ToUserID
		out[i].Status = models.StatusFriendRequest(friendRequest.Status)
		out[i].CreatedAt = friendRequest.CreateAt
	}

	return out, nil
}

func ListRequestsFromModelToDto(ctx context.Context, in models.FriendRequests) (dto.ListRequestsOUT, error) {

	requests := make([]dto.Request, len(in))

	for i, request := range in {
		requests[i] = dto.Request{
			RequestID: request.RequestID,
			Status:    dto.StatusRequest(request.Status),
		}
	}

	out := dto.ListRequestsOUT{
		Requests: requests,
	}

	return out, nil
}
