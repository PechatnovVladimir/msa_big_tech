package converter

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

// ListFriendsRequestProtoToDto - конвертация запроса из proto-формата в dto-формат
func ListFriendsRequestProtoToDto(ctx context.Context, in *social.ListFriendsRequest) (friendsIN dto.ListFriendsIN, err error) {
	if in == nil {
		return dto.ListFriendsIN{}, errors.New("grpc ListFriendsRequest is nil")
	}

	//про Cursor пока не понимаю как
	out := dto.ListFriendsIN{
		UserID: in.UserId,
		Limit:  in.Limit,
	}

	return out, nil
}

// ListFriendsResponseDtoToProto - конвертация ответа из dto в proto
func ListFriendsResponseDtoToProto(ctx context.Context, in *dto.ListFriendsOUT) (*social.ListFriendsResponse, error) {
	if in == nil {
		return nil, errors.New("grpc ListFriendsOutput is nil")
	}

	//с Cursor пока не понятно
	out := social.ListFriendsResponse{
		FriendUserIds: in.UserIDs,
		NextCursor: &social.Cursor{
			CreatedAt: nil,
		},
	}

	return &out, nil
}
