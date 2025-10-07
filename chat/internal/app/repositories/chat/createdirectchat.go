package chat

import (
	"context"
	"errors"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

// CreateDirectChat - создать личный чат
func (r *Repository) CreateDirectChat(ctx context.Context, in dtoRepo.CreateDirectChatIN) error {
	//insert в pg
	return errors.New("not implemented")
}
