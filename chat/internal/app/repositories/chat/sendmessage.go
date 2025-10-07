package chat

import (
	"context"
	"errors"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

func (r *Repository) SendMessage(ctx context.Context, in dtoRepo.SendMessageIN) (out dtoRepo.SendMessageOUT, err error) {
	return dtoRepo.SendMessageOUT{}, errors.New("not implemented")
}
