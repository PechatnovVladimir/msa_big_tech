package social

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) ListRequests(ctx context.Context, in dto.ListRequestsIN) (dto.ListRequestsOUT, error) {
	log.Println("ListRequests")

	_, err := s.AuthProvider.GetAuthUser()
	if err != nil {
		return dto.ListRequestsOUT{}, err
	}

	listRequestFromRepo, err := s.SocialRepo.ListRequests(ctx, dtoRepo.ListRequestsIN{UserID: in.UserID})
	if err != nil {
		return dto.ListRequestsOUT{}, err
	}

	listRequest, err := converter.ListRequestsFromRepoToModel(ctx, listRequestFromRepo)

	requests, err := converter.ListRequestsFromModelToDto(ctx, listRequest)
	if err != nil {
		return dto.ListRequestsOUT{}, err
	}

	return requests, nil
}
