package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) ListRequests(ctx context.Context, in dto.ListRequestsIN) (dto.ListRequestsOUT, error) {
	log.Println("ListRequests")
	return dto.ListRequestsOUT{}, nil
}
