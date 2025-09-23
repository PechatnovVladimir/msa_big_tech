package users

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"sync"
)

type RepositoryInMemory struct {
	storage map[users.UserID]users.UserProfile
	mx      sync.Mutex
}

func NewRepository(cap int) *RepositoryInMemory {
	return &RepositoryInMemory{
		storage: make(map[users.UserID]users.UserProfile, cap),
	}
}
