package social

import (
	"github.com/google/uuid"
	"time"
)

// FriendRequest - заявка в друзья
type FriendRequest struct {
	//RequestID - ID заявки
	RequestID string
	//FromUserID - ID юзера от кого заявки
	FromUserID string
	//ToUserID - ID юзера кому заявка
	ToUserID string
	//Status - статус заявки
	Status StatusFriendRequest
	//CreatedAt - время создания заявки
	CreatedAt time.Time
}

func NewFriendRequest(fromUserID string, toUserID string) *FriendRequest {
	return &FriendRequest{
		RequestID:  uuid.New().String(),
		CreatedAt:  time.Now(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Status:     PENDING,
	}
}

func (fr *FriendRequest) Accept() {
	fr.Status = ACCEPTED
}

func (fr *FriendRequest) Decline() {
	fr.Status = DECLINED
}
