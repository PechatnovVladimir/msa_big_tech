package social

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func fromListRequestsIN(in *dto.ListRequestsIN) string {
	return in.UserID
}

func toListRequestsOUT(in []*social.FriendRequest) *dto.ListRequestsOUT {
	requests := make([]*dto.Request, 0, len(in))
	for i := range in {
		requests = append(requests, &dto.Request{
			RequestID: in[i].RequestID,
			Status:    dto.StatusRequest(in[i].Status),
		})
	}

	return &dto.ListRequestsOUT{
		Requests: requests,
	}
}

//func ListRequestsFromRepoToModel(ctx context.Context, in dtoRepo.ListRequestsOUT) (social.FriendRequests, error) {
//
//	out := make(social.FriendRequests, len(in.FriendRequests))
//
//	for i, friendRequest := range in.FriendRequests {
//		out[i].RequestID = friendRequest.RequestID
//		out[i].FromUserID = friendRequest.FromUserID
//		out[i].ToUserID = friendRequest.ToUserID
//		out[i].Status = social.StatusFriendRequest(friendRequest.Status)
//		out[i].CreatedAt = friendRequest.CreateAt
//	}
//
//	return out, nil
//}
//
//func ListRequestsFromModelToDto(ctx context.Context, in social.FriendRequests) (dto.ListRequestsOUT, error) {
//
//	requests := make([]dto.Request, len(in))
//
//	for i, request := range in {
//		requests[i] = dto.Request{
//			RequestID: request.RequestID,
//			Status:    dto.StatusRequest(request.Status),
//		}
//	}
//
//	out := dto.ListRequestsOUT{
//		Requests: requests,
//	}
//
//	return out, nil
//}
