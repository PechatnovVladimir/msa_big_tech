package social

import (
	"context"
	"fmt"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) ListRequests(ctx context.Context, in *dto.ListRequestsIN) (*dto.ListRequestsOUT, error) {

	const api = "SocialService.ListRequests"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.ListRequestsOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	if in.UserID != currentUser.UserID {
		return &dto.ListRequestsOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialPermissionDenied)
	}

	data := fromListRequestsIN(in)

	listRequests, err := s.SocialRepo.ListRequests(ctx, data)
	if err != nil {
		return &dto.ListRequestsOUT{}, err
	}

	return toListRequestsOUT(listRequests), nil
}
