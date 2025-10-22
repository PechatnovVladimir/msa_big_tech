package secrets

import (
	"context"
)

type SecretsProvider interface {
	Get(ctx context.Context, key string) (string, error)      // строки (пароли, токены)
	GetBytes(ctx context.Context, key string) ([]byte, error) // бинарь (TLS ключи/серты)
}

type Secrets struct {
	SecretsProvider
}

func NewSecrets(secretsProvider SecretsProvider) *Secrets {
	return &Secrets{secretsProvider}
}
