package chat

import (
	"context"
	"errors"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

func (r *Repository) ListUserChats(ctx context.Context, in dtoRepo.ListUserChatsIN) (out dtoRepo.ListUserChatsOUT, err error) {
	return dtoRepo.ListUserChatsOUT{}, errors.New("not implemented")
}
