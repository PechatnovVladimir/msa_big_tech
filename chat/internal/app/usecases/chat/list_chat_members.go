package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func (s *Service) ListChatMembers(ctx context.Context, in *dto.ListChatMembersIN) (*dto.ListChatMembersOUT, error) {
	data := fromListChatMembersIN(in)

	userIDs, err := s.ChatRepo.ListChatMembers(ctx, data)

	if err != nil {
		return &dto.ListChatMembersOUT{}, err
	}

	out := toListChatMembersOUT(userIDs)

	return out, nil
}
