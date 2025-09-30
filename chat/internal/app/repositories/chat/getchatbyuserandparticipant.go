package chat

import (
	"context"
)

func (r *Repository) GetChatByUserAndParticipant(ctx context.Context, userID string, participantID string) (chatID string, ok bool) {
	return "", false
}
