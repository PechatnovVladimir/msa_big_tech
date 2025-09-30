package inmemory

import (
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
	"sync"
)

// InMemory - репозиторий "на коленке" исключительно для тестирования
type InMemory struct {
	userStorage  map[string]auth.User
	tokenStorage map[string]auth.Token
	mx           sync.RWMutex
}

func New() *InMemory {
	return &InMemory{
		userStorage:  make(map[string]auth.User),
		tokenStorage: make(map[string]auth.Token),
	}
}
