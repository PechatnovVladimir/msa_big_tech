package users

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"sync"
)

type InMemory struct {
	storage map[string]users.UserProfile
	mx      sync.Mutex
}

func New(cap int) *InMemory {
	return &InMemory{
		storage: make(map[string]users.UserProfile, cap),
	}
}
