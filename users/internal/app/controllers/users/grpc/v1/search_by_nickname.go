package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) SearchByNickname(ctx context.Context, request *userPB.SearchByNicknameRequest) (*userPB.SearchByNicknameResponse, error) {

	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	q := dto.Query{}

	if request.Query.IDs != nil {
		q.IDs = request.Query.IDs
	}

	if request.Query.Email != nil {
		q.Email = request.Query.Email
	}

	if request.Query.Nickname != nil {
		q.Nickname = request.Query.Nickname
	}

	if request.Query.Createdfrom != nil {
		*q.CreatedFrom = (request.Query.Createdfrom).AsTime()
	}

	if request.Query.Createdto != nil {
		*q.CreatedTo = (request.Query.Createdto).AsTime()
	}

	d := dto.SearchByNicknameDTO{
		Query: q,
		Limit: request.Limit,
	}

	p, err := s.uc.SearchByNickname(ctx, d)
	if err != nil {
		return nil, err
	}

	userProfiles := make([]*userPB.UserProfile, len(p))

	for i, _ := range p {
		userProfiles[i] = &userPB.UserProfile{
			UserId:    p[i].ID,
			Nickname:  p[i].Nickname,
			Email:     p[i].Email,
			AvatarUrl: p[i].Avatar,
			Bio:       p[i].Bio,
		}
	}

	return &userPB.SearchByNicknameResponse{UserProfile: userProfiles}, nil
}
