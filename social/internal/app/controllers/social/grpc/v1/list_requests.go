package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func (s *Service) ListRequests(ctx context.Context, request *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	data := fromListRequestsRequestToDto(request)

	requests, err := s.SocialUseCase.ListRequests(ctx, data)
	if err != nil {
		return nil, err
	}

	return fromDtoToListRequestsResponse(requests), nil
}
