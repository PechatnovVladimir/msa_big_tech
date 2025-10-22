package secrets

import (
	"context"
)

type VaultProvider struct {
}

func (v *VaultProvider) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (v *VaultProvider) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func NewVaultProvider() *VaultProvider {
	return &VaultProvider{}
}
