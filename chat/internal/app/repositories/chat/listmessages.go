package chat

import (
	"context"
	"errors"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

func (r *Repository) ListMessages(ctx context.Context, in dtoRepo.ListMessagesIN) (out dtoRepo.ListMessagesOUT, err error) {
	return dtoRepo.ListMessagesOUT{}, errors.New("not implemented")
}
