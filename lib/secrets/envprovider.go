package secrets

import (
	"context"
	"errors"
	"os"
)

type EnvProvider struct {
}

func (env *EnvProvider) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (env *EnvProvider) Get(ctx context.Context, key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("key not found")
	}

	return value, nil
}

func NewEnvProvider() *EnvProvider {
	return &EnvProvider{}
}
