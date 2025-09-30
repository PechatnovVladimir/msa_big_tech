package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) ListRequests(ctx context.Context, request *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	log.Println("SocialService ListRequests called")
	return nil, nil
}
