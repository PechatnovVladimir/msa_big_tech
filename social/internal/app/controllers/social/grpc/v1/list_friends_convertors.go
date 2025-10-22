package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func fromListFriendsRequestToDto(in *social.ListFriendsRequest) *dto.ListFriendsIN {
	cursor := &dto.Cursor{
		ID:   in.Cursor.UserId,
		Time: in.Cursor.CreatedAt.AsTime(),
	}

	return &dto.ListFriendsIN{
		UserID: in.UserId,
		Limit:  in.Limit,
		Cursor: *cursor,
	}
}

func fromDtoToListFriendsResponse(in *dto.ListFriendsOUT) *social.ListFriendsResponse {
	if in == nil {
		return nil
	}
	nextCursor := &social.Cursor{
		UserId:    in.Cursor.ID,
		CreatedAt: timestamppb.New(in.Cursor.Time),
	}

	friendIDs := make([]string, len(in.UserIDs))
	for i := range in.UserIDs {
		friendIDs[i] = in.UserIDs[i]
	}

	return &social.ListFriendsResponse{
		NextCursor:    nextCursor,
		FriendUserIds: friendIDs,
	}
}
