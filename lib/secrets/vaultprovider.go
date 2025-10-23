package secrets

import (
	"context"
)

type VaultProvider struct {
	isSet bool
}

func (v *VaultProvider) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (v *VaultProvider) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (v *VaultProvider) IsSet(ctx context.Context) bool {
	return v.isSet
}

func NewVaultProvider() *VaultProvider {
	return &VaultProvider{
		isSet: true,
	}
}
