package secrets

import (
	"errors"
	"os"
)

type EnvProvider struct {
	isSet bool
}

func (env *EnvProvider) GetBytes(key string) ([]byte, error) {
	return nil, nil
}

func (env *EnvProvider) Get(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("key not found")
	}

	return value, nil
}

func (env *EnvProvider) IsSet() bool {
	return env.isSet
}

func NewEnvProvider() *EnvProvider {
	return &EnvProvider{
		isSet: true,
	}
}
