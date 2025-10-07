package models

import (
	"github.com/google/uuid"
	"time"
)

// FriendRequest - заявка в друзья
type FriendRequest struct {
	//RequestID - ID заявки
	RequestID string `json:"request_id"`
	//FromUserID - ID юзера от кого заявки
	FromUserID string `json:"from_user_id"`
	//ToUserID - ID юзера кому заявка
	ToUserID string `json:"to_user_id"`
	//Status - статус заявки
	Status StatusFriendRequest `json:"status"`
	//CreatedAt - время создания заявки
	CreatedAt time.Time `json:"created_at"`
}

type FriendRequests []FriendRequest

func NewFriendRequest(fromUserID string, toUserID string) FriendRequest {
	return FriendRequest{
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
