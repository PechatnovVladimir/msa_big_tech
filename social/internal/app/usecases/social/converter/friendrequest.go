package converter

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
)

func FriendRequestsFromRepoToModel(ctx context.Context, in dtoRepo.FriendRequest) (models.FriendRequest, error) {

	return models.FriendRequest{
		RequestID:  in.RequestID,
		FromUserID: in.FromUserID,
		ToUserID:   in.ToUserID,
		Status:     models.StatusFriendRequest(in.Status),
		CreatedAt:  in.CreateAt,
	}, nil
}

func FriendRequestFromModelToRepo(ctx context.Context, in models.FriendRequest) (dtoRepo.FriendRequest, error) {
	return dtoRepo.FriendRequest{
		RequestID:  in.RequestID,
		FromUserID: in.FromUserID,
		ToUserID:   in.ToUserID,
		Status:     int(in.Status),
		CreateAt:   in.CreatedAt,
	}, nil
}
