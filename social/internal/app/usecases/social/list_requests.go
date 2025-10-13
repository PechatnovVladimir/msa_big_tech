package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) ListRequests(ctx context.Context, in *dto.ListRequestsIN) (*dto.ListRequestsOUT, error) {

	data := fromListRequestsIN(in)

	listRequests, err := s.SocialRepo.ListRequests(ctx, data)
	if err != nil {
		return &dto.ListRequestsOUT{}, err
	}

	return toListRequestsOUT(listRequests), nil
}
