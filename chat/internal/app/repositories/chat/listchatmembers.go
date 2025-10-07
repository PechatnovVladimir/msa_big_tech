package chat

import (
	"context"
	"errors"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

func (r *Repository) ListChatMembers(ctx context.Context, in dtoRepo.ListChatMembersIN) (out dtoRepo.ListChatMembersOUT, err error) {
	return out, errors.New("not implemented")
}
